package search

import (
	"log"
	"testing"

	"github.com/sebas7dk/go-recipes/config"
	"github.com/sebas7dk/go-recipes/models"
	"github.com/stretchr/testify/require"
)

var (
	v = map[string]interface{}{
		"title":        "Fruit Recipe",
		"ingredients":  "apple, pear, kiwi, banana",
		"instructions": "Cut the fruit and put it in a bowl.",
		"time":         10,
		"people":       2,
	}
	c        *Connection
	recipeId string
)

func TestMain(m *testing.M) {
	setup()
	//run all the tests
	m.Run()
	//clean the index
	teardown()
}

func setup() {
	var err error

	if err := config.NewConfig("../.env"); err != nil {
		log.Fatal("Error loading the .env file")
	}

	//Use the test index
	SetIndex(config.Get("ES_TEST_INDEX"))

	c, err = NewConnection()
	if err != nil {
		log.Fatal(err)
	}

}

func teardown() {
	if _, err := c.DeleteIndex(); err != nil {
		log.Fatal("Error unable to delete the index")
	}
}

func TestNewConnection(t *testing.T) {
	c, err := NewConnection()

	require.Nil(t, err)
	require.NotNil(t, c)
}

func TestCreate(t *testing.T) {
	r := models.Recipe{
		Title:        v["title"].(string),
		Time:         v["time"].(int),
		People:       v["people"].(int),
		Ingredients:  v["ingredients"].(string),
		Instructions: v["instructions"].(string),
	}

	o, err := c.Create(r)
	require.Nil(t, err)

	recipeId = o.Id
}

func TestShow(t *testing.T) {
	o, err := c.Show()

	require.Nil(t, err)
	require.NotNil(t, o)

	CheckRecipes(o, t)
}

func TestGetById(t *testing.T) {
	o, err := c.GetById(recipeId)

	require.Nil(t, err)
	require.NotNil(t, o)

	require.Equal(t, recipeId, o.Id)
	require.Equal(t, v["title"], o.Title)
	require.Equal(t, v["time"], o.Time)
	require.Equal(t, v["people"], o.People)
	require.Equal(t, v["ingredients"], o.Ingredients)
	require.Equal(t, v["instructions"], o.Instructions)
}

func TestUpdate(t *testing.T) {
	newTitle := "Updated title"

	r := models.Recipe{
		Title: newTitle,
	}

	o, err := c.Update(recipeId, r)

	require.Nil(t, err)
	require.NotNil(t, o)

	out, err := c.GetById(recipeId)

	require.Nil(t, err)
	require.NotNil(t, o)

	require.Equal(t, newTitle, out.Title)
}

func TestQuery(t *testing.T) {
	o, err := c.Query(v["title"].(string))

	require.Nil(t, err)
	require.NotNil(t, o)

	CheckRecipes(o, t)
}

func TestDelete(t *testing.T) {
	o, err := c.Delete(recipeId)

	require.Nil(t, err)
	require.NotNil(t, o)

	_, err = c.GetById(recipeId)

	require.NotNil(t, err)
}

func CheckRecipes(recipes []models.Recipe, t *testing.T) {
	for _, r := range recipes {
		require.Equal(t, recipeId, r.Id)
		require.Equal(t, v["title"], r.Title)
		require.Equal(t, v["time"], r.Time)
		require.Equal(t, v["people"], r.People)
		require.Equal(t, v["ingredients"], r.Ingredients)
		require.Equal(t, v["instructions"], r.Instructions)
	}
}
