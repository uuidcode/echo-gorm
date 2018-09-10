package book

import (
	"github.com/echo-gorm/database"
	"github.com/echo-gorm/model/page"
	"github.com/echo-gorm/model/user"
	"github.com/echo-gorm/util"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"net/http"
	"time"
)

type Book struct {
	BookId      int64     `gorm:"PRIMARY_KEY" json:"bookId" form:"bookId" query:"bookId"`
	UserId      int64     `json:"userId" form:"userId" query:"userId"`
	Name        string    `json:"name" form:"name" query:"name"`
	RegDatetime time.Time `json:"regDatetime" form:"regDatetime" query:"regDatetime"`
	ModDatetime time.Time `json:"modDatetime" form:"modDatetime" query:"modDatetime"`
}

func (Book) TableName() string {
	return "book"
}

func Index(c echo.Context) error {
	var bookList []Book
	book := new(Book)
	err := c.Bind(book)
	util.CheckErr(err)

	var totalCount int64
	database.MainDB.Model(book).Count(&totalCount)

	p := page.NewWithContext(c, totalCount)
	database.MainDB.Offset(p.ItemOffset).Limit(p.ItemLimit).Find(&bookList)

	c.Logger().Debug(coreutil.ToJson(bookList))

	pageUrl := util.GetUrlAndRemovePageParam(c.Request())

	return c.Render(http.StatusOK, "book/index", echo.Map{
		"bookList": bookList,
		"page":     p,
		"url":      pageUrl,
	})
}

func Form(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	result := echo.Map{}

	if book.BookId != 0 {
		database.MainDB.First(book, Book{
			BookId: book.BookId,
		})

		result["book"] = book
	}

	return c.Render(http.StatusOK, "book/form.html", result)
}

func Post(c echo.Context) error {
	book := new(Book)
	user := new(user.User)
	err := c.Bind(book)
	util.CheckErr(err)

	tx := database.Begin(database.MainDB)

	user.RegDatetime = time.Now()
	user.ModDatetime = time.Now()
	user.Name = util.CreateUuid()
	database.Save(tx, &user)

	c.Logger().Debug(">>> user:", util.ToJson(user))

	book.RegDatetime = time.Now()
	book.ModDatetime = time.Now()
	book.UserId = user.UserId

	database.Save(tx, &book)
	database.Commit(tx)

	c.Logger().Debug(util.ToJson(book))

	return c.JSON(http.StatusOK, book)
}

func Put(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	newBook := new(Book)

	database.MainDB.First(newBook, Book{
		BookId: book.BookId,
	})

	newBook.Name = book.Name

	database.MainDB.Save(newBook)
	return c.JSON(http.StatusOK, book)
}

func Delete(c echo.Context) error {
	book := new(Book)
	err := c.Bind(book)
	coreutil.CheckErr(err)

	database.MainDB.Delete(&book)

	return c.JSON(http.StatusOK, book.BookId)
}

func Get(c echo.Context) error {
	bookIdValue := c.Param("bookId")

	book := new(Book)
	bookId, err := coreutil.ParseInt(bookIdValue)
	coreutil.CheckErr(err)

	database.MainDB.First(book, Book{
		BookId: bookId,
	})

	return c.JSON(http.StatusOK, book)
}
