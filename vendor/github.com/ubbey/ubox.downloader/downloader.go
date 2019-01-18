/*
	desc: 下载器定义
	author: licowei
	time: 2018-12-22
	ext:
*/

/*
	每个任务使用一个协程进行下载，在http头部加入rang字段实现任务的开始、暂停、断点续传等。
	任务协程采用channel来控制
*/

package downloader

import (
	"fmt"
	"github.com/ubbey/ubox.downloader/model"
	"sync"
)

type Downloader struct {
	tasklist map[string]*worker
	mut      sync.Mutex
}

var downloaderIns *Downloader

func GetDownloader() *Downloader {
	if downloaderIns == nil {
		downloaderIns = new(Downloader)
		downloaderIns.tasklist = make(map[string]*worker)
	}
	return downloaderIns
}

//func init() {
//	RebuildCache()
//}

func RebuildCache(business string) {
	tsDo := model.DownloadTask{}
	tslist, _ := tsDo.QueryAllDownloadTask(business)

	GetDownloader().mut.Lock()
	defer GetDownloader().mut.Unlock()

	for _, v := range tslist {

		if _, ok := GetDownloader().tasklist[v.Taskid]; ok {
			continue
		}

		w := newWorkerWithTask(v)

		if w.task.Status == DOWNLOAD_RUNNING {
			// 该行不能优化删除
			w.task.Status = DOWNLOAD_INIT
			w.Start()
		}
		GetDownloader().tasklist[v.Taskid] = w
	}
}

func (dl *Downloader) NewTaskWork(id string) (status int, err error) {

	if id == "" {
		return 0, fmt.Errorf("new work failed , task id null")
	}

	tsDo := model.DownloadTask{Taskid: id}
	ts, _ := tsDo.QueryDownloadTaskByid()
	if ts.Taskid == "" {
		return 0, fmt.Errorf("new work failed , task not found in db")
	}

	dl.mut.Lock()
	defer dl.mut.Unlock()

	if _, ok := dl.tasklist[id]; ok {
		if dl.tasklist[id].GetTaskStatus() == DOWNLOAD_RUNNING {
			return DOWNLOAD_RUNNING, nil
		}
	}

	w := newWorkerWithTask(ts)

	w.Start()

	dl.tasklist[id] = w

	return DOWNLOAD_RUNNING, nil
}

func (dl *Downloader) GetTaskList() (tsl []*model.DownloadTask) {
	dl.mut.Lock()
	defer dl.mut.Unlock()

	for _, v := range dl.tasklist {
		tsl = append(tsl, v.CopyTask())
	}

	return
}

func (dl *Downloader) GetTaskByid(id string) (ts *model.DownloadTask) {
	dl.mut.Lock()
	defer dl.mut.Unlock()

	if _, ok := dl.tasklist[id]; ok {
		return dl.tasklist[id].CopyTask()
	}

	// cache miss , load from db
	tsdo := model.DownloadTask{}
	tsdo.Taskid = id
	tsdo, _ = tsdo.QueryDownloadTaskByid()

	// db hit , update cache
	if tsdo.Taskid == id {
		w := newWorkerWithTask(tsdo)
		if w.task.Status == DOWNLOAD_RUNNING {
			w.task.Status = DOWNLOAD_INIT
			w.Start()
		}
		dl.tasklist[id] = w
		return &tsdo
	}

	return nil
}

func (dl *Downloader) ChangTaskStatus(id string, status int) error {
	dl.mut.Lock()
	defer dl.mut.Unlock()

	if _, ok := dl.tasklist[id]; ok {
		switch status {
		case ACTION_START:
			dl.tasklist[id].Start()
		case ACTION_PAUSE:
			dl.tasklist[id].Pause()
		default:
			return fmt.Errorf("action invaild")
		}
		return nil
	}
	return fmt.Errorf("task id not found")
}

func (dl *Downloader) DelTask(id string) error {
	dl.mut.Lock()
	defer dl.mut.Unlock()

	if _, ok := dl.tasklist[id]; ok {
		dl.tasklist[id].Stop()
		tsdo := model.DownloadTask{}
		tsdo.Taskid = id
		err := tsdo.RemoveDownloadTaskByid()
		return err
	}
	return fmt.Errorf("task id not found")
}
