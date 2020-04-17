package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ellvisca/convert-node-go/app"
	"github.com/ellvisca/convert-node-go/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//User router
	router.HandleFunc("/api/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/me", controllers.CurrentUser).Methods("GET")

	//Task router
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/me", controllers.MyTask).Methods("GET")

	//Middleware
	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println("Listening on port ", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
