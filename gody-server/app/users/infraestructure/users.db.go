package infraestructure

import (
	"time"

	usersModel "github.com/gody-server/app/users/domain/models"
	// uncomment this a change _ for db when you are making database queries
	bcrypt "github.com/gody-server/adapters/bcrypt"
	db "github.com/gody-server/adapters/database"
)

type UsersDb struct {
	// Add any dependencies or configurations related to the UserRepository here if needed.
}

func NewUsersDb() *UsersDb {
	// Initialize any dependencies and configurations for the UsersRepository here if needed.
	return &UsersDb{}
}

func (u *UsersDb) CreateUsers(users usersModel.UsersCreate) error {
	// Insert into database new user
	user := usersModel.Users{Username: users.Username, Password: users.Password, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	result := db.Client().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	// Implement your creation logic here
	return nil
}

func (u *UsersDb) GetUsers() ([]usersModel.Users, error) {
	// Implement your retrieval logic here
	var users []usersModel.Users
	users = append(users, usersModel.Users{ID: 1, Username: "username", Password: "password", CreatedAt: time.Time{}, UpdatedAt: time.Time{}})
	return users, nil
}

func (u *UsersDb) GetOneUsers(usersId int) (usersModel.Users, error) {
	// Implement your single retrieval logic here
	return usersModel.Users{}, nil
}

func (u UsersDb) UpdateUsers(usersId int, user usersModel.UsersPut) error {
	hashedPassword, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result := db.Client().Model(usersModel.Users{}).Where("id = ?", usersId).Updates(usersModel.Users{Username: user.Username, Password: hashedPassword, UpdatedAt: time.Time{}})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UsersDb) DeleteUsers(usersId int) error {
	// Implement your deletion logic here
	return nil
}

func (u UsersDb) Login(users usersModel.LoginRequest) (usersModel.Users, error) {
	// select user from database
	var user usersModel.Users
	result := db.Client().Where("username = ?", users.Username).First(&user)
	if result.Error != nil {
		return usersModel.Users{}, result.Error
	}
	return user, nil
}
