package services

import (
	"errors"

	jwt "github.com/gody-server/adapters/jwt"
	tokenModel "github.com/gody-server/app/token/domain/models"
	tokenInterface "github.com/gody-server/app/token/domain/repositories"
)

type Service struct {
	tokenRepository tokenInterface.IToken
}

func NewService(tokenRepository tokenInterface.IToken) *Service {
	return &Service{
		tokenRepository: tokenRepository,
	}
}

func (s *Service) CreateToken(token tokenModel.Token) error {
	err := s.tokenRepository.CreateToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetToken() ([]tokenModel.Token, error) {
	result, err := s.tokenRepository.GetToken()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetOneToken(tokenId int) (tokenModel.Token, error) {
	result, err := s.tokenRepository.GetOneToken(tokenId)
	if err != nil {
		return tokenModel.Token{}, err
	}
	return result, nil
}

func (s *Service) UpdateToken(tokenId int) error {
	err := s.tokenRepository.UpdateToken(tokenId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteToken(tokenId int) error {
	err := s.tokenRepository.DeleteToken(tokenId)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) Verify(token tokenModel.TokenVerify) error {
	result := jwt.ValidateToken(token.Token)
	if result == "Error" {
		return errors.New("Error")
	}
	// err := s.tokenRepository.Verify(token)
	// if err != nil {
	// 	return err
	// }
	return nil
}
