package router

import (
	"net/http"

	"github.com/go-recipes/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"GET",
		"/recipes/{id}",
		handlers.ShowRecipeById,
	},
	Route{
		"GET",
		"/recipes",
		handlers.ShowAll,
	},
	Route{
		"GET",
		"/recipes/search/{term}",
		handlers.SearchRecipe,
	},
	Route{
		"POST",
		"/recipes",
		handlers.CreateRecipe,
	},
	Route{
		"PUT",
		"/recipes/{id}",
		handlers.UpdateRecipeById,
	},
	Route{
		"DELETE",
		"/recipes/{id}",
		handlers.DestroyRecipeById,
	},
}
