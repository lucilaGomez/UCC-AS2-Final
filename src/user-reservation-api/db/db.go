package db

import (
	"user-reservation-api/client"
	"user-reservation-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	DBName := "hotel-reservation"
	DBUser := "root"
	DBPass := "pass"
	DBHost := "mysql"

	Db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Add all clients here
	client.Db = Db

}

func StartDbEngine() {
	// Migrate all model classes
	Db.AutoMigrate(&model.User{})

	log.Info("Finishing Migration Database Tables")
}
