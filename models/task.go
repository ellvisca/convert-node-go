package models

import (
	u "github.com/ellvisca/convert-node-go/utils"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title      string `json:"title"`
	DueDate    string `json:"dueDate"`
	Importance int    `json:"importance"`
	Completed  bool   `json:"false"`
	UserId     uint   `json:"user_id"`
}

func (task *Task) Create() map[string]interface{} {

	GetDB().Create(task)

	if task.ID <= 0 {
		return u.Message(false, "Failed to create task, connection error")
	}

	response := u.Message(true, "Task has been created")
	response["task"] = task
	return response
}
