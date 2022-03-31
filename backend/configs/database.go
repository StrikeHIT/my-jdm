package configs

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Open() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/myjdm?charset=utf8&parseTime=True")

	if err != nil {
		log.Panic(err)
	}

	return db
}

func Close() error {
	return db.Close()
}
