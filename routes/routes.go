package routes

import (
	"go-backend/handlers"
	"go-backend/models"
)

type Route = models.Route

func GetRoutes() []Route {
	var routes []Route
	routes = append(routes, Route{Route: "/", Handler: handlers.Home})
	routes = append(routes, Route{Route: "/count", Handler: handlers.GetUserCount})
	routes = append(routes, Route{Route: "/users", Handler: handlers.GetUsers})
	return routes
}
