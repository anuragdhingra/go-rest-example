package db

import (
	"github.com/anuragdhingra/go-rest-example/db/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB

func init() {

	log.Println(getDatasource())
	conn, err := gorm.Open("mysql", "monstar-lab:password@tcp(mysql:3306)/testapplication?parseTime=true")
	if err != nil {
		panic("Error connecting to DB")
	}

	db = conn
	db.AutoMigrate(models.User{})
}

func GetDb() *gorm.DB  {
	return db
}

func getDatasource() (dataSource string) {
	dataSource = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") +
		"@tcp(mysql:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true"
	return
}
