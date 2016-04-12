package search

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-recipes/config"
	"github.com/go-recipes/models"
	elastigo "github.com/mattbaird/elastigo/lib"
)

type Connection struct {
	Conn *elastigo.Conn
}

var index string

//SetIndex set the index name
func SetIndex(i string) {
	index = i
}

//NewConnection create a new Elastic Search connection
func NewConnection() (*Connection, error) {
	c := elastigo.NewConn()
	if c == nil {
		return nil, errors.New("Unable to connect to Elastic Search")
	}

	c.Domain = config.Get("ES_DOMAIN")
	c.Port = config.Get("ES_PORT")

	conn := &Connection{Conn: c}

	return conn, nil
}

//Show all the docs in the index
func (c *Connection) Show() ([]models.Recipe, error) {
	searchJSON := `{
      "query" : {
          "match_all" : {}
      }
  }`

	o, err := c.Conn.Search(index, "recipe", nil, searchJSON)
	r := BuildResults(o.Hits.Hits)

	return r, err
}

//GetById show the doc by id
func (c *Connection) GetById(id string) (*models.Recipe, error) {
	var recipe *models.Recipe

	o, err := c.Conn.Get(index, "recipe", id, nil)

	if err == nil {
		json.Unmarshal(*o.Source, &recipe)
		recipe.Id = o.Id
	}

	return recipe, err
}

//Create a new doc
func (c *Connection) Create(r models.Recipe) (elastigo.BaseResponse, error) {
	return c.Conn.Index(index, "recipe", "", nil, r)
}

//Update a doc by id
func (c *Connection) Update(id string, r models.Recipe) (elastigo.BaseResponse, error) {
	return c.Conn.Index(index, "recipe", id, nil, r)
}

//Query the index and match the search term
func (c *Connection) Query(s string) ([]models.Recipe, error) {
	searchJSON := fmt.Sprintf(`{
	    "query" : {
	        "multi_match": {
	            "query" : "%s",
	            "fields" : ["title^50", "category^30", "instructions^25", "ingredients^20"]
	        }
	    }
	}`, s)

	o, err := c.Conn.Search(index, "recipe", nil, searchJSON)
	r := BuildResults(o.Hits.Hits)

	return r, err
}

//Delete a doc from the index
func (c *Connection) Delete(id string) (elastigo.BaseResponse, error) {
	return c.Conn.Delete(index, "recipe", id, nil)
}

//DeleteIndex alll docs from the index
func (c *Connection) DeleteIndex() (elastigo.BaseResponse, error) {
	return c.Conn.DeleteIndex(index)
}

//BuildResults loop through the hits based on the total hits
func BuildResults(recipes []elastigo.Hit) []models.Recipe {
	var recipe models.Recipe
	rs := make(models.Recipes, 0)

	for _, r := range recipes {
		if err := json.Unmarshal(*r.Source, &recipe); err == nil {
			recipe.Id = r.Id
			rs = append(rs, recipe)
		}
	}

	return rs
}
