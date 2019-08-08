package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hipercubus/jazmeendeco/config"
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB
// var err error

// User user
type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}

// func connect() {
// 	db, err = gorm.Open("mysql", "root:123@/jazmeendeco?parseTime=true")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

// InitialMigration migracion
func InitialMigration() {
	db := config.Connect()
	defer db.Close()

	db.AutoMigrate(&User{})
}

// AllUsers route
func AllUsers(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// OneUser route
func OneUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	db.Where("id = ?", id).Find(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// NewUser route
func NewUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Created")
}

// DeleteUser route
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()
	fmt.Fprintf(w, "Delete User Endpoint")
}

// UpdateUser route
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint")
}
