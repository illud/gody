package services

import (
	bcrypt "github.com/gody-server/adapters/bcrypt"
	jwt "github.com/gody-server/adapters/jwt"
	usersModel "github.com/gody-server/app/users/domain/models"
	usersInterface "github.com/gody-server/app/users/domain/repositories"
)

type Service struct {
	usersRepository usersInterface.IUsers
}

func NewService(usersRepository usersInterface.IUsers) *Service {
	return &Service{
		usersRepository: usersRepository,
	}
}

func (s *Service) CreateUsers(users usersModel.UsersCreate) error {
	err := s.usersRepository.CreateUsers(users)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUsers() ([]usersModel.Users, error) {
	result, err := s.usersRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetOneUsers(usersId int) (usersModel.Users, error) {
	result, err := s.usersRepository.GetOneUsers(usersId)
	if err != nil {
		return usersModel.Users{}, err
	}
	return result, nil
}

func (s *Service) UpdateUsers(usersId int, user usersModel.UsersPut) error {
	err := s.usersRepository.UpdateUsers(usersId, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUsers(usersId int) error {
	err := s.usersRepository.DeleteUsers(usersId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Login(users usersModel.LoginRequest) (string, error) {

	result, err := s.usersRepository.Login(users)
	if err != nil {
		return "", err
	}

	// check password
	checkPassword := bcrypt.CheckPasswordHash(users.Password, result.Password)
	if !checkPassword {
		return "", err
	}

	// generate token
	token := jwt.GenerateToken(users.Username)
	if token == "" {
		return "", err
	}

	return token, nil
}
