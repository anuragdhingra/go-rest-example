package db

import (
	"github.com/anuragdhingra/go-rest-example/db/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	)

var db *gorm.DB

func init() {
	conn, err := gorm.Open("mysql", "monstar-lab:password@/testapplication?charset=utf8&parseTime=True")

	if err != nil {
		panic("Error connecting to DB")
	}
	db = conn
	db.AutoMigrate(models.User{})
}

func GetDb() *gorm.DB  {
	return db
}