package main

import (
	"log"
	"net/http"

	"github.com/go-recipes/config"
	"github.com/go-recipes/handlers"
	"github.com/go-recipes/router"
	"github.com/go-recipes/search"
	"github.com/rs/cors"
)

func main() {
	if err := config.NewConfig(".env"); err != nil {
		log.Fatal("Error loading the .env file")
	}

	c, err := search.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	search.SetIndex(config.Get("ES_INDEX"))
	//Set the elastigo connection
	handlers.SetConnection(c)

	//Create a new router
	r := router.NewRouter()

	//Allow cors requests
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	log.Fatal(http.ListenAndServe(config.Get("APP_PORT"), handler))
}
