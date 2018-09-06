package context

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type WebContext struct {
	echo.Context
	DB *gorm.DB
}

func GetWebContext(c echo.Context) *WebContext {
	return c.(*WebContext)
}
