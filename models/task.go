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
	Completed        bool               `json:"completed"`
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

//Get current user tasks
func MyTask(userId primitive.ObjectID) []*Task {
	tasks := make([]*Task, 0)

	filter := bson.M{"owner": userId}
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
		// decode similar to deserialize process.
		err := cur.Decode(&task)
		if err != nil {
			fmt.Println(err)
		}

		// add item our array
		tasks = append(tasks, task)
	}

	return tasks
}

//Edit user task
func (task *Task) Edit(userId, taskId primitive.ObjectID) map[string]interface{} {
	filter := bson.M{"_id": taskId, "owner": userId}
	collection := GetDB().Collection("tasks")

	update := bson.M{
		"$set": bson.M{
			"title":      task.Title,
			"dueDate":    task.DueDate,
			"importance": task.Importance,
			"completed":  task.Completed,
		},
	}

	collection.UpdateOne(context.TODO(), filter, update)

	response := u.Message(true, "Task has been edited")
	response["task"] = task
	return response
}
