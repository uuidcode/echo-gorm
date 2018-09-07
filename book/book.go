package book

import (
	"github.com/echo-gorm/database"
	"github.com/echo-gorm/util"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"net/http"
	"time"
)

type Book struct {
	BookId      int64     `gorm:"PRIMARY_KEY" json:"bookId" form:"bookId" query:"bookId"`
	UserId      int64     `json:"bookId" form:"userId" query:"userId"`
	Name        string    `json:"name" form:"name" query:"name"`
	RegDatetime time.Time `json:"regDatetime" form:"regDatetime" query:"regDatetime"`
	ModDatetime time.Time `json:"modDatetime" form:"modDatetime" query:"modDatetime"`
}

func (Book) TableName() string {
	return "book"
}

func Index(c echo.Context) error {
	var bookList []Book

	offset, limit := util.GetOffsetAndLimit(c)
	database.DB.Offset(offset).Limit(limit).Find(&bookList)

	c.Logger().Debug(coreutil.ToJson(bookList))

	return c.Render(http.StatusOK, "book/index", echo.Map{
		"bookList": bookList,
	})
}

func Form(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	result := echo.Map{}

	if book.BookId != 0 {
		database.DB.First(book, Book{
			BookId: book.BookId,
		})

		result["book"] = book
	}

	return c.Render(http.StatusOK, "book/form.html", result)
}

func Post(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	book.RegDatetime = time.Now()
	book.ModDatetime = time.Now()
	book.UserId = 1
	database.DB.Create(&book)

	return c.JSON(http.StatusOK, book)
}

func Put(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	newBook := new(Book)

	database.DB.First(newBook, Book{
		BookId: book.BookId,
	})

	newBook.Name = book.Name

	database.DB.Save(newBook)
	return c.JSON(http.StatusOK, book)
}

func Delete(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	database.DB.Delete(&book)

	return c.JSON(http.StatusOK, book.BookId)
}

func Get(c echo.Context) error {
	bookIdValue := c.Param("bookId")

	book := new(Book)
	bookId, err := coreutil.ParseInt(bookIdValue)
	coreutil.CheckErr(err)

	database.DB.First(book, Book{
		BookId: bookId,
	})

	return c.JSON(http.StatusOK, book)
}
