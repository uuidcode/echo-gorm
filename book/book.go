package book

import (
	"github.com/echo-gorm/context"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	BookId      int64 `gorm:"PRIMARY_KEY" json:"bookId" form:"bookId" query:"bookId"`
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
	webContext := context.GetWebContext(c)
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	result := echo.Map{}

	if book.BookId != 0 {
		webContext.DB.First(book, Book{
			BookId: book.BookId,
		})

		result["book"] = book
	}

	return c.Render(http.StatusOK, "book/form.html", result)
}

func Post(c echo.Context) error {
	webContext := context.GetWebContext(c)

	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	book.RegDatetime = time.Now()
	book.ModDatetime = time.Now()
	book.UserId = 1
	webContext.DB.Create(&book)

	return c.JSON(http.StatusOK, book)
}

func Put(c echo.Context) error {
	webContext := context.GetWebContext(c)

	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	newBook := new(Book)

	webContext.DB.First(newBook, Book{
		BookId: book.BookId,
	})

	newBook.Name = book.Name

	webContext.DB.Save(newBook)
	return c.JSON(http.StatusOK, book)
}

func Remove(c echo.Context) error {
	webContext := context.GetWebContext(c)

	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	webContext.DB.Delete(&book, Book{
		BookId: book.BookId,
	})

	return c.JSON(http.StatusOK, book.BookId)
}

func Get(c echo.Context) error {
	bookIdValue := c.Param("bookId")
	webContext := context.GetWebContext(c)

	book := new(Book)
	bookId, err := strconv.ParseInt(bookIdValue, 10, 64)
	coreutil.CheckErr(err)

	webContext.DB.First(book, Book{
		BookId: bookId,
	})

	return c.JSON(http.StatusOK, book)
}
