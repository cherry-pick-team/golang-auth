package Router

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang-auth/src/Handlers"
	"golang-auth/src/Utils"
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
		"VkAuthRedirect",
		"GET",
		"/auth/vk",
		Handlers.VkAuthRedirect,
	},
	Route{
		"TodoIndex",
		"GET",
		"/auth/vk/callback",
		Handlers.VkAuthCode,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
