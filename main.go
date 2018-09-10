package main

import (
	"github.com/echo-gorm/app"
	"github.com/echo-gorm/middleware"
	"github.com/echo-gorm/model/book"
	"github.com/echo-gorm/model/user"
)

func main() {
	e := app.Echo()

	e.GET("/book", book.Index)
	e.GET("/book/:bookId", book.Get)
	e.GET("/book/form", book.Form)
	e.POST("/book", book.Post)
	e.PUT("/book", book.Put)
	e.DELETE("/book", book.Delete)

	e.GET("/user", user.Index)
	e.GET("/user/:userId", user.Get)
	e.GET("/user/form", user.Form)
	e.POST("/user", user.Post)
	e.PUT("/user", user.Put)
	e.DELETE("/user", user.Delete)

	e.Use(middleware.Hello)
	e.Start(":17004")
}
