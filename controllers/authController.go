package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ellvisca/convert-node-go/models"
	u "github.com/ellvisca/convert-node-go/utils"
)

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
