package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/uuidcode/coreutil"
)

const url = "root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true"

var DB, err = gorm.Open("mysql", url)

func init() {
	coreutil.CheckErr(err)
	DB.LogMode(true)
}
