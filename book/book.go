package book

import (
	"time"
	"net/http"
	"github.com/labstack/echo"
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
	return context.Render(http.StatusOK, "book/index", echo.Map{

	})
}
