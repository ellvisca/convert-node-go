package models

import (
	"fmt"

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

//Create task
func (task *Task) Create() map[string]interface{} {

	GetDB().Create(task)

	if task.ID <= 0 {
		return u.Message(false, "Failed to create task, connection error")
	}

	response := u.Message(true, "Task has been created")
	response["task"] = task
	return response
}

//Get current user task
func MyTask(user uint) []*Task {
	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Find(&tasks).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return tasks
}
