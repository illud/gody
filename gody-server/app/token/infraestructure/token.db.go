package infraestructure

import (
	tokenModel "github.com/gody-server/app/token/domain/models"
	// uncomment this a change _ for db when you are making database queries
	_ "github.com/gody-server/adapters/database"
)

type TokenDb struct {
	// Add any dependencies or configurations related to the UserRepository here if needed.
}

func NewTokenDb() *TokenDb {
	// Initialize any dependencies and configurations for the TokenRepository here if needed.
	return &TokenDb{}
}

var token []tokenModel.Token

func (t *TokenDb) CreateToken(token tokenModel.Token) error {
	// Implement your creation logic here
	return nil
}

func (t *TokenDb) GetToken() ([]tokenModel.Token, error) {
	// Implement your retrieval logic here
	var token []tokenModel.Token
	token = append(token, tokenModel.Token{Id: 1})
	return token, nil
}

func (t *TokenDb) GetOneToken(tokenId int) (tokenModel.Token, error) {
	// Implement your single retrieval logic here
	return tokenModel.Token{Id: tokenId}, nil
}

func (t TokenDb) UpdateToken(tokenId int) error {
	// Implement your update logic here
	return nil
}

func (t TokenDb) DeleteToken(tokenId int) error {
	// Implement your deletion logic here
	return nil
}
func (t TokenDb) Verify(token tokenModel.TokenVerify) error {
	// Implement your creation logic here
	return nil
}
