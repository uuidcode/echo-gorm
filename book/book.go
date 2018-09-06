package book

import (
	"encoding/json"
	"github.com/echo-gorm/context"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	BookId      int64 `gorm:"PRIMARY_KEY"`
	UserId      int64
	Name        string
	RegDatetime time.Time
	ModDatetime time.Time
}

func (Book) TableName() string {
	return "book"
}

func Index(c echo.Context) error {
	webContext := context.GetWebContext(c)
	var bookList []Book

	webContext.DB.Find(&bookList)

	return c.Render(http.StatusOK, "book/index", echo.Map{
		"bookList": bookList,
	})
}

func Form(c echo.Context) error {
	return c.Render(http.StatusOK, "book/form.html", echo.Map{})
}

func getBook(webContext *context.WebContext) Book {
	b, err := ioutil.ReadAll(webContext.Request().Body)
	coreutil.CheckErr(err)

	webContext.Logger().Debug("book", string(b))

	var book Book
	err = json.Unmarshal(b, &book)
	coreutil.CheckErr(err)

	return book
}

func Save(c echo.Context) error {
	webContext := context.GetWebContext(c)
	book := getBook(webContext)
	book.RegDatetime = time.Now()
	book.ModDatetime = time.Now()
	book.UserId = 1
	webContext.DB.Create(&book)

	return c.JSON(http.StatusOK, book)
}

func Remove(c echo.Context) error {
	webContext := context.GetWebContext(c)
	book := getBook(webContext)

	webContext.DB.Delete(&book, Book{
		BookId: book.BookId,
	})

	return c.JSON(http.StatusOK, book.BookId)
}

func Get(c echo.Context) error {
	bookIdValue := c.Param("bookId")
	webContext := context.GetWebContext(c)

	var book Book
	bookId, err := strconv.ParseInt(bookIdValue, 10, 64)
	coreutil.CheckErr(err)

	webContext.DB.First(&book, Book{
		BookId: bookId,
	})

	return c.JSON(http.StatusOK, &book)
}
