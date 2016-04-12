package router

import (
	"github.com/go-recipes/config"
	"github.com/gorilla/mux"
)

//NewRouter build the mux router
func NewRouter() *mux.Router {

	router := mux.NewRouter().
		StrictSlash(true).
		PathPrefix("/api/" + config.Get("APP_VERSION")).
		Subrouter()

	for _, route := range routes {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	}

	return router
}
