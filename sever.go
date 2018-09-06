package main

import (
	"fmt"
	"github.com/echo-gorm/book"
	"github.com/echo-gorm/context"
	"github.com/foolin/echo-template"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/uuidcode/coreutil"
)

func main() {
	url := "root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)
	db.LogMode(true)
	defer db.Close()

	e := echo.New()
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.WebContext{c, db}
			return h(cc)
		}
	})

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println(string(reqBody))
	}))

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
	e.POST("/book", book.Save)
	e.GET("/book/form", book.Form)

	e.Start(":17004")
}
