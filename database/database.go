package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	config "qr-menu-project-backend/config"
)

var DB *gorm.DB

func NewDB(Params...string) *gorm.DB{
	var err error

	conString := config.GetPostgresConnectionString()


	log.Print(conString)
	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	
	if err!= nil {
        log.Fatal(err)
    }
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}