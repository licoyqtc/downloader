package main

import (
	"github.com/labstack/echo"
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

	e.Start(":10086")
}
