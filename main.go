package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hipercubus/jazmeendeco/models"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "JazmeenDeco - Rest API")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", models.AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", models.OneUser).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", models.NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}", models.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", models.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("GO is listening")
	models.InitialMigration()
	handleRequest()
}
