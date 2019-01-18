/*
	desc: worker定义
	author: licowei
	time: 2018-12-23
	ext:
*/

package downloader

import (
	"encoding/json"
	"github.com/ubbey/ubox.downloader/model"
	"log"
	"os"
	"sync"
)

type http_arg struct {
	Header   map[string][]string
	Method   string
	Body     interface{}
	Off_size int64
}

func newHttpArgs() *http_arg {
	args := &http_arg{}
	args.Header = make(map[string][]string)
	args.Off_size = default_size
	return args
}

type worker struct {
	mut  sync.Mutex
	task *model.DownloadTask
	args *http_arg
	stop chan struct{}
}

func newWorker() *worker {
	w := new(worker)
	w.args = newHttpArgs()
	return w
}

func newWorkerWithTask(ts model.DownloadTask) *worker {
	w := newWorker()
	json.Unmarshal([]byte(ts.Header), &w.args.Header)
	json.Unmarshal([]byte(ts.Body), &w.args.Body)
	w.args.Method = ts.Method
	w.PasteTask(&ts)
	return w
}

func (w *worker) Start() {
	w.mut.Lock()
	defer w.mut.Unlock()

	if w.task.Status == DOWNLOAD_RUNNING {
		return
	}
	w.task.Status = DOWNLOAD_RUNNING

	go func() {
		for {
			// 拷贝对象，减小锁粒度
			tmpTask := w.CopyTask()

			var try int64 = 0
			select {
			case <-w.stop:
				return
			default:
				var n int64
				var packsize int64
				var err error
				if tmpTask.Status == DOWNLOAD_RUNNING && tmpTask.CurOffset != 0 {
					n, packsize, err = download(tmpTask, w.args, os.O_WRONLY|os.O_APPEND)
				} else {
					n, packsize, err = download(tmpTask, w.args, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
				}

				// 下载失败错误处理
				if err != nil {
					try++
					if try == 10 {
						log.Println("task id ", tmpTask.Taskid, " down load failed : ", err.Error())
						w.Cancel()
					} else {
						log.Println("down load err : ", err.Error())
					}
					break
				}
				// 成功，计数清零
				try = 0

				tmpTask.Finish += n
				tmpTask.CurOffset += n
				tmpTask.Total = packsize

				if tmpTask.CurOffset >= tmpTask.Total {
					log.Println("task : ", tmpTask.Name, " download success : ", tmpTask.Finish, tmpTask.Total)
					w.PasteTask(tmpTask)
					w.Stop()
					break
				}

				_, err = tmpTask.UpdateDownloadTaskByid()
				if err != nil {
					log.Println("update task status err , id ", tmpTask.Taskid, " err : ", err.Error())
					break
				}

				// db update success , update cache
				w.PasteTask(tmpTask)
			}
			//time.Sleep(time.Second)
		}
	}()
}

func (w *worker) Pause() {
	w.ChangeStatus(DOWNLOAD_PAUSE)
}

func (w *worker) Cancel() {
	w.ChangeStatus(DOWNLOAD_FAILED)
}

func (w *worker) Stop() {
	w.ChangeStatus(DOWNLOAD_FINISH)
}

func (w *worker) GetTaskStatus() int {
	w.mut.Lock()
	status := w.task.Status
	w.mut.Unlock()

	return status
}

func (w *worker) CopyTask() *model.DownloadTask {
	var tmpTask model.DownloadTask
	w.mut.Lock()
	tmpTask = *w.task
	w.mut.Unlock()
	return &tmpTask
}

func (w *worker) PasteTask(task *model.DownloadTask) {
	var tmpTask model.DownloadTask = *task
	w.mut.Lock()
	w.task = &tmpTask
	w.mut.Unlock()
}

func (w *worker) ChangeStatus(status int) {
	tmpTask := w.CopyTask()
	if tmpTask.Status == status {
		return
	}

	tmpTask.Status = status
	_, err := tmpTask.UpdateDownloadTaskByid()
	if err != nil {
		return
	}

	stop := false
	// db update success , update cache
	w.mut.Lock()
	if w.task.Status == DOWNLOAD_RUNNING {
		stop = true
	}
	w.task.Status = status
	w.mut.Unlock()

	if stop {
		w.stop <- struct{}{}
	}
}
