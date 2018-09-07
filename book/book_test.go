package book

import (
	"fmt"
	"github.com/echo-gorm/app"
	"github.com/echo-gorm/database"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestBook(t *testing.T) {
	var book Book
	database.DB.DropTable(&book)
	database.DB.CreateTable(&book)
}

func TestGet(t *testing.T) {
	e := app.TestEcho()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/book/:bookId")
	c.SetParamNames("bookId")
	c.SetParamValues("3")

	Get(c)

	bytes, err := ioutil.ReadAll(rec.Body)
	coreutil.CheckErr(err)

	fmt.Println(string(bytes))
}

func TestForm(t *testing.T) {
	e := app.TestEcho()
	req := httptest.NewRequest(echo.GET, "/book/form?bookId=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	Form(c)

	fmt.Println(rec.Body.String())
}
