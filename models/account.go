package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (acc *Account) Create() map[string]interface{} {

}

func (acc *Account) Validate() (map[string]interface{}, bool) {

}

func GetUser(u uint) *Account {

}
