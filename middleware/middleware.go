package middleware

import (
	"github.com/labstack/echo"
	"path"
	"runtime"
)

func Hello(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pc, file, line, ok := runtime.Caller(0)

		if ok {
			funcName := runtime.FuncForPC(pc).Name()
			c.Logger().Debugf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
		}

		return next(c)
	}
}
