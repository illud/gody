package repositories

import (
	tokenModel "github.com/gody-server/app/token/domain/models"
)

type IToken interface {
	CreateToken(token tokenModel.Token) error
	GetToken() ([]tokenModel.Token, error)
	GetOneToken(tokenId int) (tokenModel.Token, error)
	UpdateToken(tokenId int) error
	DeleteToken(tokenId int) error
	Verify(token tokenModel.TokenVerify) error
}
