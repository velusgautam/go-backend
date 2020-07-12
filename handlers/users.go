package handlers

import (
	"encoding/json"
	"go-backend/models"
	"net/http"
)

// User model from models
type User = models.User

// UserCount model from models
type UserCount = models.UserCount

var users []User

// GetUsers handler to fetch all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users = append(users, User{Name: "John Doe", Email: "johndoe@gmail.com"})
	users = append(users, User{Name: "Kim Jane", Email: "kimjane@gmail.com"})
	users = append(users, User{Name: "Philip Mathew", Email: "pmathew@gmail.com"})
	// setting content-type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	return
}

// GetUserCount handler to fetch the users count
func GetUserCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count := UserCount{Count: len(users)}
	json.NewEncoder(w).Encode(count)
	return
}
