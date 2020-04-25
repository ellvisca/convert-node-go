package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ellvisca/todolist/models"
	u "github.com/ellvisca/todolist/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CreateTask = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(primitive.ObjectID)
	task := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	task.Owner = userId
	resp := task.Create()
	u.Respond(w, resp)
}

var MyTask = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(primitive.ObjectID)
	data := models.MyTask(userId)

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var EditTask = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(primitive.ObjectID)
	keys := r.URL.Query()["taskId"]
	taskId, _ := primitive.ObjectIDFromHex(keys[0])

	task := &models.Task{}
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := task.Edit(userId, taskId)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteTask = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(primitive.ObjectID)
	keys := r.URL.Query()["taskId"]
	taskId, _ := primitive.ObjectIDFromHex(keys[0])

	task := &models.Task{}
	resp := task.Delete(userId, taskId)
	u.Respond(w, resp)
}
