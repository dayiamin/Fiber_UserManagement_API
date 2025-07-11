package database

import (
	"log"
	"github.com/dayiamin/Fiber_UserManagement_API/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DataBase *gorm.DB


// ConnectToDb initializes the SQLite database connection,
// performs auto-migration for models, and assigns the connection to the global DataBase variable.
//
// If any error occurs during connection or migration, the function logs the error and exits the application.
func ConnectToDb(){

	db , err := gorm.Open(sqlite.Open("data.db"),&gorm.Config{})
	if err !=nil{
		log.Fatal("Error in database Connection ",err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil{
		log.Fatal("Migration error ",err)
	}

	DataBase = db

}