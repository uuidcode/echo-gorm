package user

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

func TestUser(t *testing.T) {
	var user User
	database.MainDB.DropTable(&user)
	database.MainDB.CreateTable(&user)
}

func TestGet(t *testing.T) {
	e := app.TestEcho()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:userId")
	c.SetParamNames("userId")
	c.SetParamValues("3")

	Get(c)

	bytes, err := ioutil.ReadAll(rec.Body)
	coreutil.CheckErr(err)

	fmt.Println(string(bytes))
}

func TestForm(t *testing.T) {
	e := app.TestEcho()
	req := httptest.NewRequest(echo.GET, "/user/form?userId=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	Form(c)

	fmt.Println(rec.Body.String())
}
