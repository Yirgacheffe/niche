package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ReturnConfig",
		"GET",
		"/configs",
		ConfigHandler,
	},
	Route{
		"HealthCheck",
		"GET",
		"/configs/health",
		HealthCheckHandler,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).Handler(handler)
	}

	return router

}
