package models

import "net/http"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCount struct {
	Count int `json:"count"`
}

type Route struct {
	Route   string
	Handler func(http.ResponseWriter, *http.Request)
}
