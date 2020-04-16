package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ellvisca/convert-node-go/models"
	u "github.com/ellvisca/convert-node-go/utils"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	task := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	task.UserId = user
	resp := task.Create()
	u.Respond(w, resp)
}
