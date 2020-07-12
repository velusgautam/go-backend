package handlers

import (
	"fmt"
	"net/http"
)

// Home handler give the default page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go REST SERVICE")
}
