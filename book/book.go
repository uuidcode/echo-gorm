package book

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Book struct {
	BookId      int
	UserId      int
	Name        string
	RegDatetime time.Time
	ModDatetime time.Time
}

func (Book) TableName() string {
	return "book"
}

func Index(context echo.Context) error {
	return context.Render(http.StatusOK, "book/index", echo.Map{})
}

func Form(context echo.Context) error {
	return context.Render(http.StatusOK, "book/form.html", echo.Map{})
}
