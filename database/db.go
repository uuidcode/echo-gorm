package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/uuidcode/coreutil"
)

const url = "root:rootroot@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=true"

var MainDB, err = gorm.Open("mysql", url)

func init() {
	coreutil.CheckErr(err)
	MainDB.LogMode(true)
}

func Begin(currentDB *gorm.DB) *gorm.DB {
	db := currentDB.Begin()

	if db.Error != nil {
		panic(db.Error)
	}

	return db
}

func Save(currentDB *gorm.DB, model interface{}) *gorm.DB {
	db := currentDB.Save(model)

	if db.Error != nil {
		currentDB.Rollback()
		panic(db.Error)
	}

	return db
}

func Commit(currentDB *gorm.DB) {
	err = currentDB.Commit().Error

	if err != nil {
		currentDB.Rollback()
		panic(err)
	}
}
