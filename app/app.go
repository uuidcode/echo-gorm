package app

import (
	"fmt"
	"github.com/echo-gorm/database"
	"github.com/echo-gorm/logger"
	"github.com/foolin/echo-template"
	"github.com/labstack/echo"
	"github.com/plutov/echo-logrus"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"net/http"
	"os"
	"strings"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	uri := c.Request().RequestURI
	c.Logger().Debugf(">>> uri: %v", uri)

	if strings.HasSuffix(uri, ".map") {
		return
	}

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

func TestEcho() *echo.Echo {
	os.Chdir("..")
	return Echo()
}

func Echo() *echo.Echo {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&prefixed.TextFormatter{})
	logger := logrus.StandardLogger()

	database.DB.SetLogger(gomlogger.New())

	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler

	echoLogger := echologrus.Logger{logger}
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

	logger.Debug("Echo is ready")

	return e
}
