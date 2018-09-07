package main

import (
	"github.com/echo-gorm/app"
	"github.com/echo-gorm/book"
)

func main() {
	e := app.Echo()
	e.GET("/book", book.Index)
	e.GET("/book/:bookId", book.Get)
	e.GET("/book/form", book.Form)
	e.POST("/book", book.Post)
	e.PUT("/book", book.Put)
	e.DELETE("/book", book.Delete)

	e.Start(":17004")
}
