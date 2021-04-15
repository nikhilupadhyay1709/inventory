package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func Init() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "./inventory.db")
	if err != nil {
		panic(err)
	}
}
