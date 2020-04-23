package models

import (
	"context"
	"os"

	"github.com/Kamva/mgm"
	"github.com/dgrijalva/jwt-go"
	u "github.com/ellvisca/todolist/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Token            string `json:"token"`
}

type Token struct {
	UserId primitive.ObjectID
	jwt.StandardClaims
}

//Create new user
func (user *User) Create() map[string]interface{} {
	encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(encrypted_password)

	collection := GetDB().Collection("users")
	collection.InsertOne(context.TODO(), user)

	response := u.Message(true, "User has been created")
	response["user"] = user
	return response
}

func Login(email, password string) map[string]interface{} {
	user := &User{}

	filter := bson.M{"email": email}
	collection := GetDB().Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection, please try again")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	response := u.Message(true, "Logged in")
	response["user"] = user
	return response
}

//Get current user
func Current(u primitive.ObjectID) *User {
	user := &User{}

	filter := bson.M{"_id": u}
	collection := GetDB().Collection("users")
	collection.FindOne(context.TODO(), filter).Decode(&user)

	if user.Email == "" {
		return nil
	}

	return user
}
