package repositories

import (
	usersModel "github.com/gody-server/app/users/domain/models"
)

type IUsers interface {
	CreateUsers(users usersModel.UsersCreate) error
	GetUsers() ([]usersModel.Users, error)
	GetOneUsers(usersId int) (usersModel.Users, error)
	UpdateUsers(usersId int, user usersModel.UsersPut) error
	DeleteUsers(usersId int) error
	Login(users usersModel.LoginRequest) (usersModel.Users, error)
}
