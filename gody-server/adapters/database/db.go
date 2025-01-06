package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	// DB The database connection
	db *gorm.DB
)

// Connect to database
func Connect() {
	//CONNECTION
	dbCon, err := gorm.Open(sqlite.Open("gody.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR: ", err)
	}

	// Migrate the schema
	// dbCon.AutoMigrate(&tasksModel.Task{})

	// defer db.Close()
	db = dbCon
	fmt.Println("CONNECTED")
}

func Client() *gorm.DB {
	return db
}
