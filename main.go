package main

import (
	"fmt"
	"go-backend/routes"
	"log"
	"net/http"
)

func handleRequest() {
	for _, v := range routes.GetRoutes() {
		http.HandleFunc(v.Route, v.Handler)
	}
	fmt.Println("Serving from : http://localhost:2001")
	log.Fatal(http.ListenAndServe(":2001", nil))
}

func main() {
	handleRequest()
}
