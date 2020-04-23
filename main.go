package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ellvisca/todolist/app"
	"github.com/ellvisca/todolist/controllers"
	"github.com/gorilla/mux"
	"github.com/maple-ai/syrup"
)

func main() {
	router := syrup.New(mux.NewRouter())

	//User router
	router.Post("/api/user", controllers.CreateUser)
	router.Post("/api/user/login", controllers.Authenticate)
	router.Get("/api/user/me", controllers.CurrentUser)

	//Task router
	router.Post("/api/task", controllers.CreateTask)
	router.Get("/api/task/me", controllers.MyTask)

	//Middleware
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Listening on port ", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
