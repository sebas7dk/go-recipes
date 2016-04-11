package main

import (
	"log"
	"net/http"

	"github.com/go-recipes/config"
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
	defer search.Conn.Close()

	search.SetIndex(config.ENV["ES_INDEX"])

	r := router.NewRouter()

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)

	log.Fatal(http.ListenAndServe(config.ENV["APP_PORT"], handler))
}
