package models

import (
	u "innohack-backend/utils"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId uuid.uuid
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (acc *Account) Create() map[string]interface{} {
	if res, ok := acc.Validate(); !ok {
		return res
	}

	passHash, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost))
	acc.Password := string(passHash)

	 GetDB().Create(acc)

	if acc.ID == "" {
		return u.Message(false, "Failed to create account. Incorrect id")
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{UserId: acc.ID})
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	acc.Token = tokenString
	acc.Password = ""

	res := u.Message(true, "Account created")
	res["account"] = acc
	return res
}

func (acc *Account) Validate() (map[string]interface{}, bool) {
	if len(acc.Login) < 5 {
		return u.Message(false, "Password must be more than 5 symbols"), false
	}

	if len(acc.Password) < 6 {
		return u.Message(false, "Password must be more than 6 symbols"), false
	}

	tmpAcc := &Account{}

	err := GetDB().Table("accounts").Where("login = ?", acc.Login).First(tmpAcc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Login failed"), false
	}

	return u.Message(false, "Validate passed"), true
}

func GetUser(u uint) *Account {

}
