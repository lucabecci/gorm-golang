package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/lucabecci/gorm-golang/src/models"
)

//Data is a struct for handling errors and configs
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

	//set the value of the db in Task module
	models.InstanceDatabase(data.Db)
	return data
}

//GetDatabase is a function for the connection of the db
func GetDatabase() *gorm.DB {
	db := database
	return db
}

//CreateTables is a func for the Migration of the tables in gorm
func CreateTables(db *gorm.DB) {
	exists := db.Migrator().HasTable(&models.Task{})
	if exists == false {
		db.AutoMigrate(&models.Task{})
		fmt.Println("Models Created")
		return
	}
	return

}
