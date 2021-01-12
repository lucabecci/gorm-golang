package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/lucabecci/gorm-golang/src/models"
)

type Data struct {
	Db      *gorm.DB
	Success bool
	Error   error
}

var database *gorm.DB

//GetConnection is a function for the connection of the db
func GetConnection() Data {
	var data Data
	dsn, _ := os.LookupEnv("DB_URI")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), //logger blocked
	})

	if err != nil {
		data.Db = &gorm.DB{}
		data.Success = false
		data.Error = err
		return data
	}

	data.Db = db
	data.Success = true
	database = data.Db

	//migrations of the schemas
	CreateTables(data.Db)
	return data
}

//GetConnection is a function for the connection of the db
func GetDatabase() (*gorm.DB, bool) {
	db := database
	return db, true
}

func CreateTables(db *gorm.DB) {
	exists := db.Migrator().HasTable(&models.Task{})
	if exists == false {
		db.AutoMigrate(&models.Task{})
		fmt.Println("Models Created")
		return
	}
	return
}
