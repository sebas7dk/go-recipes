package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sebas7dk/go-recipes/config"
	"github.com/sebas7dk/go-recipes/models"
	"github.com/sebas7dk/go-recipes/search"
	"github.com/gorilla/mux"
)

type jsonError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var c *search.Connection

func SetConnection(conn *search.Connection) {
	c = conn
}

//Index show the welcome message
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Go Recipe "+config.Get("APP_VERSION"))
}

//ShowAll show all recipes
func ShowAll(w http.ResponseWriter, r *http.Request) {
	if r, err := c.Show(); err != nil {
		BuildResponse(w, "not_found")
	} else {
		BuildResponse(w, "ok")
		if err := json.NewEncoder(w).Encode(r); err != nil {
			log.Fatal(err)
		}
	}
}

//SearchRecipe search the recipe
func SearchRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if r, err := c.Query(string(vars["term"])); err != nil {
		BuildResponse(w, "not_found")
	} else {
		BuildResponse(w, "ok")
		if err := json.NewEncoder(w).Encode(r); err != nil {
			log.Fatal(err)
		}
	}
}

//ShowRecipeById find the recipe by id
func ShowRecipeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if r, err := c.GetById(string(vars["id"])); err != nil {
		BuildResponse(w, "not_found")
	} else {
		BuildResponse(w, "ok")
		if err := json.NewEncoder(w).Encode(r); err != nil {
			log.Fatal(err)
		}
	}
}

//CreateRecipe Create a new recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe

	if err := json.Unmarshal(ValidJsonBody(w, r), &recipe); err != nil {
		BuildResponse(w, "unprocessed")
		return
	}

	if _, err := c.Create(recipe); err != nil {
		fmt.Println(err)
		BuildResponse(w, "bad_request")
	} else {
		BuildResponse(w, "created")
	}
}

//UpdateRecipeById perform an update on the doc
func UpdateRecipeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var recipe models.Recipe

	if err := json.Unmarshal(ValidJsonBody(w, r), &recipe); err != nil {
		BuildResponse(w, "unprocessed")
		return
	}

	if _, err := c.Update(vars["id"], recipe); err != nil {
		BuildResponse(w, "bad_request")
	} else {
		BuildResponse(w, "ok")
	}
}

//DestroyRecipeById remove the doc by id
func DestroyRecipeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, err := c.Delete(vars["id"]); err != nil {
		BuildResponse(w, "not_found")
	} else {
		BuildResponse(w, "ok")
	}
}

//ValidJsonBody validate the json body and that it doesn't exceed the limit
func ValidJsonBody(w http.ResponseWriter, r *http.Request) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}

	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}

	return body
}

//BuildResponse Build the headers and return a json encoded error if needed
func BuildResponse(w http.ResponseWriter, s string) {
	var c int
	var t string
	var e bool

	switch s {
	case "not_found":
		c = http.StatusNotFound
		t = "Recipe not found"
		e = true
	case "bad_request":
		c = http.StatusBadRequest
		t = "Unable to create recipe, field is missing"
		e = true
	case "unprocessed":
		c = 422
		t = "Unable to process the json body"
		e = true
	case "ok":
		c = http.StatusOK
		e = false
	case "created":
		c = http.StatusCreated
		e = false
	default:
		log.Fatal("unrecognized status")
	}

	// return valid json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8,")
	// set the correct header with the status code
	w.WriteHeader(c)

	if e {
		if err := json.NewEncoder(w).Encode(jsonError{Code: c, Message: t}); err != nil {
			log.Fatal(err)
		}
	}
}
