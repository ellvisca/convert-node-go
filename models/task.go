package models

import (
	"context"
	"fmt"

	"github.com/Kamva/mgm"
	u "github.com/ellvisca/todolist/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string             `json:"title"`
	DueDate          string             `json:"dueDate"`
	Importance       int                `json:"importance"`
	Completed        bool               `json:"false"`
	Owner            primitive.ObjectID `json:"owner"`
}

//Create task
func (task *Task) Create() map[string]interface{} {
	collection := GetDB().Collection("tasks")
	collection.InsertOne(context.TODO(), task)

	response := u.Message(true, "Task has been created")
	response["task"] = task
	return response
}

//Get current user task
func MyTask(user primitive.ObjectID) []*Task {
	tasks := make([]*Task, 0)

	filter := bson.M{"owner": user}
	collection := GetDB().Collection("tasks")
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		task := &Task{}
		// & character returns the memory address of the following variable.
		err := cur.Decode(&task) // decode similar to deserialize process.
		if err != nil {
			fmt.Println(err)
		}

		// add item our array
		tasks = append(tasks, task)
	}

	return tasks
}
