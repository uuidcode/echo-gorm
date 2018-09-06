package main

import (
	"fmt"
	"github.com/echo-gorm/book"
	"github.com/echo-gorm/context"
	"github.com/echo-gorm/logger"
	"github.com/foolin/echo-template"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/plutov/echo-logrus"
	"github.com/sirupsen/logrus"
	"github.com/uuidcode/coreutil"
	"net/http"
	"os"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	errorPage := fmt.Sprintf("%d.html", code)

	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}

	c.Logger().Error(err)
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger := logrus.New()

	url := "root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)

	db.LogMode(true)
	db.SetLogger(gomlogger.New())
	defer db.Close()

	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.WebContext{c, db}
			return h(cc)
		}
	})

	echoLogger := echologrus.Logger{logrusLogger}
	e.Logger = echoLogger
	e.Use(echologrus.Hook())

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
	e.DELETE("/book", book.Remove)
	e.GET("/book/form", book.Form)

	e.Start(":17004")
}
