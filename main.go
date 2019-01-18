package main

import (
	"fmt"
	"github.com/ubbey/ubox.downloader"
	"github.com/ubbey/ubox.downloader/model"
)

type ss struct {
	S []string    `json:"s"`
	I []int       `json:"i"`
	A interface{} `json:"a"`
}

func main() {

	tsDo := model.DownloadTask{}
	tsDo.Business = downloader.BUSINESS_UBOX
	tsDo.Subbusiness = "demo"
	tsDo.User = "test"
	tsDo.Url = "http://iamtest.yqtc.co/searchData/pvr.png"
	tsDo.Header = ""
	tsDo.Method = "GET"
	tsDo.Body = ""
	tsDo.Dir = "/tmp/packets" // todo 容器内进程运行目录+req.dir
	tsDo.Name = "get_pvr.png"

	tsDo.Generateid()

	downloader.EnsureDirExist(tsDo.Dir)

	_, err := tsDo.NewDownloadTask()
	if err != nil {
		fmt.Println("get fail : ", err.Error())
	}

	_, err = downloader.GetDownloader().NewTaskWork(tsDo.Taskid)
	if err != nil {
		fmt.Println("new task fail : ", err.Error())
	}

	tsDo = model.DownloadTask{}
	tsDo.Business = downloader.BUSINESS_UBOX
	tsDo.Subbusiness = "demo"
	tsDo.User = "test"
	tsDo.Url = "http://iamtest.yqtc.co/searchData/pvr.png"
	tsDo.Header = ""
	tsDo.Method = "POST"
	tsDo.Body = ""
	tsDo.Dir = "/tmp/packets" // todo 容器内进程运行目录+req.dir
	tsDo.Name = "post_pvr.png"

	tsDo.Generateid()

	downloader.EnsureDirExist(tsDo.Dir)

	_, err = tsDo.NewDownloadTask()
	if err != nil {
		fmt.Println("get fail : ", err.Error())
	}

	_, err = downloader.GetDownloader().NewTaskWork(tsDo.Taskid)
	if err != nil {
		fmt.Println("new task fail : ", err.Error())
	}
}
