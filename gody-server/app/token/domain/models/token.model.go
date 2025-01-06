package models

type Token struct {
	Id int
}

type TokenVerify struct {
	Token string `json:"token"`
}
