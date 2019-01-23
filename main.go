package main

import (
	"encoding/json"
	"github.com/yqtc.com/ubox.uapp/uvm/sdk"
	"github.com/yqtc.com/ubox.uapp/uvm/sdk/echo"
	"log"
)

type ss struct {
	S []string    `json:"s"`
	I []int       `json:"i"`
	A interface{} `json:"a"`
}

func main() {

	e := echo.New()

	e.POST("/uapp/com.pvr.www/greet", func(context echo.Context) error {
		return context.JSONPretty(200, "hi i am post pvr...\n", "")
	})
	e.GET("/uapp/com.pvr.www/greet", func(context echo.Context) error {
		return context.JSONPretty(200, "hi i am get pvr...\n", "")
	})

	e.POST("/uapp/com.pvr.www/download", func(context echo.Context) error {
		req := sdk.Sdk_downloader_task_download_req{}
		req.Business = sdk.UAPP_PVR
		req.Subbusiness = "test"
		req.Method = "GET"
		req.Dir = "/usr/local/application/com.pvr.www/workroot/data"
		req.Url = "http://iamtest.yqtc.co/Application/com.pvr.www/pvr-1.0.0.xml"
		req.Packname = "pvr-1.0.0.xml"

		rsp, err := sdk.Sdk_downloader_task_download(req)
		if err != nil {
			log.Println(err.Error())
			return context.JSONPretty(200, err.Error()+"\n", "")
		}

		data, _ := json.Marshal(rsp)
		return context.JSONPretty(200, string(data)+"\n", "")
	})

	if true {
		err := e.Start("unix:///var/run/echo.sock")
		log.Println("Unreachable code, error:%s", err.Error())
	} else {
		err := e.Start(":8087")
		log.Println("Unreachable code, error:%s", err.Error())
	}
}
