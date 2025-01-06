package database

import (
	"time"

	bcrypt "github.com/gody-server/adapters/bcrypt"
	actionsModel "github.com/gody-server/app/actions/domain/models"
	usersModel "github.com/gody-server/app/users/domain/models"
)

func Migrate() error {

	Client().AutoMigrate(&usersModel.Users{}, &actionsModel.Actions{})
	// Inserting a user (this will create a new row in the Users table)

	// Check if the Users table already has any data
	var count int64
	Client().Model(&usersModel.Users{}).Count(&count)

	if count == 0 {
		hashedPassword, err := bcrypt.HashPassword("toor")
		if err != nil {
			return err
		}

		defaultUser := usersModel.Users{
			Username:  "root",
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Use Create method to insert data
		Client().Create(&defaultUser)
	}

	return nil
}
