package book

import (
	"github.com/echo-gorm/context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/uuidcode/coreutil"
	"net/http/httptest"
	"testing"
)

func TestBook(t *testing.T) {
	url := "root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)

	db.LogMode(true)

	defer db.Close()

	var book Book
	db.DropTable(&book)
	db.CreateTable(&book)
}

func TestPost(t *testing.T) {
	url := "root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)
	db.LogMode(true)
	defer db.Close()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/book/:bookId")
	c.SetParamNames("bookId")
	c.SetParamValues("1")

	cc := &context.WebContext{c, db}

	Get(cc)
}
