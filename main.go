package main

import (
	"log"
	"yqtc.com/ubox.uapp/uvm/sdk/echo"
)

type ss struct {
	S []string    `json:"s"`
	I []int       `json:"i"`
	A interface{} `json:"a"`
}

func main() {

	e := echo.New()

	e.POST("/uapp/com.pvr.www/greet", func(context echo.Context) error {
		return context.JSONPretty(200, "hi i am post pvr...", "")
	})
	e.GET("/uapp/com.pvr.www/greet", func(context echo.Context) error {
		return context.JSONPretty(200, "hi i am get pvr...", "")
	})

	if true {
		err := e.Start("unix:///var/run/echo.sock")
		log.Println("Unreachable code, error:%s", err.Error())
	} else {
		err := e.Start(":8087")
		log.Println("Unreachable code, error:%s", err.Error())
	}
}
