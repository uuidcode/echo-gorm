package book

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/uuidcode/coreutil"
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
