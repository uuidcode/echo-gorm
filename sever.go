package main

import (
	"github.com/echo-gorm/book"
	"github.com/foolin/echo-template"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	templateConfig := echotemplate.TemplateConfig{
		Root:         "template",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	}
	e.Renderer = echotemplate.New(templateConfig)

	e.Static("/static", "static")
	e.File("/favicon.ico", "static/ico/favicon.ico")
	e.GET("/book", book.Index)

	e.Start(":17004")
}
