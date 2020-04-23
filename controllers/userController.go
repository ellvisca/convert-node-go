package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ellvisca/todolist/models"
	u "github.com/ellvisca/todolist/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = models.GetDB().Collection("users")

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

var CurrentUser = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(primitive.ObjectID)
	data := models.Current((user))

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
