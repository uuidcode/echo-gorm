package book

import (
	"fmt"
	"github.com/echo-gorm/app"
	"github.com/echo-gorm/database"
	"github.com/echo-gorm/util"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestBook(t *testing.T) {
	var book Book
	database.MainDB.DropTable(&book)
	database.MainDB.CreateTable(&book)
}

func TestInsert(t *testing.T) {
	for i := 0; i < 601; i++ {
		database.MainDB.Create(&Book{
			Name:        util.CreateUuid(),
			RegDatetime: time.Now(),
			ModDatetime: time.Now(),
			UserId:      1,
		})
	}
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

func TestPost(t *testing.T) {
	e := app.TestEcho()
	json := `{"name":"Hello"}`

	req := httptest.NewRequest(echo.POST, "/book", strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	Post(c)

	fmt.Println(rec.Body.String())
}
