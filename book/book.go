package book

import (
	"encoding/json"
	"fmt"
	"github.com/echo-gorm/context"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"io/ioutil"
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

func (book *Book) Save(db *gorm.DB) {
	db.Create(book)
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

func Save(c echo.Context) error {
	webContext := context.GetWebContext(c)

	b, err := ioutil.ReadAll(webContext.Request().Body)
	coreutil.CheckErr(err)

	fmt.Println("Save", string(b))

	var book Book
	err = json.Unmarshal(b, &book)
	coreutil.CheckErr(err)

	book.RegDatetime = time.Now()
	book.ModDatetime = time.Now()
	book.UserId = 1
	webContext.DB.Create(&book)

	return c.JSON(http.StatusOK, book)
}
