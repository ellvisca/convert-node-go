package models

import (
	u "github.com/ellvisca/convert-node-go/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Create() map[string]interface{} {
	encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(encrypted_password)

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error")
	}

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}
