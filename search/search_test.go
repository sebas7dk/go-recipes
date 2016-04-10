package search

import (
	"log"
	"testing"

	"github.com/go-recipes/config"
	"github.com/go-recipes/models"
	"github.com/stretchr/testify/require"
)

var (
	StrValues = map[string]string{
		"title":        "Fruit Recipe",
		"ingredients":  "apple, pear, kiwi, banana",
		"instructions": "Cut the fruit and put it in a bowl.",
	}
	IntValues = map[string]int{
		"time":   10,
		"people": 2,
	}
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
	if err := config.NewConfig("../.env"); err != nil {
		log.Fatal("Error loading the .env file")
	}

	//Use the test index
	SetIndex(config.ENV["ES_TEST_INDEX"])
}

func teardown() {
	if _, err := DeleteIndex(); err != nil {
		log.Fatal("Error unable to delete the index")
	}
}

func TestNewConnection(t *testing.T) {
	c := Connect()

	require.NotNil(t, c)
}

func TestCreate(t *testing.T) {
	r := models.Recipe{
		Title:        StrValues["title"],
		Time:         IntValues["time"],
		People:       IntValues["people"],
		Ingredients:  StrValues["ingredients"],
		Instructions: StrValues["instructions"],
	}

	o, err := Create(r)
	require.Nil(t, err)

	recipeId = o.Id
}

func TestShow(t *testing.T) {
	o, err := Show()

	require.Nil(t, err)
	require.NotNil(t, o)

	CheckRecipes(o, t)
}

func TestGetById(t *testing.T) {
	o, err := GetById(recipeId)

	require.Nil(t, err)
	require.NotNil(t, o)

	require.Equal(t, recipeId, o.Id)
	require.Equal(t, StrValues["title"], o.Title)
	require.Equal(t, IntValues["time"], o.Time)
	require.Equal(t, IntValues["people"], o.People)
	require.Equal(t, StrValues["ingredients"], o.Ingredients)
	require.Equal(t, StrValues["instructions"], o.Instructions)
}

func TestUpdate(t *testing.T) {
	newTitle := "Updated title"

	r := models.Recipe{
		Title: newTitle,
	}

	o, err := Update(recipeId, r)

	require.Nil(t, err)
	require.NotNil(t, o)

	out, err := GetById(recipeId)

	require.Nil(t, err)
	require.NotNil(t, o)

	require.Equal(t, newTitle, out.Title)
}

func TestQuery(t *testing.T) {
	o, err := Query(StrValues["title"])

	require.Nil(t, err)
	require.NotNil(t, o)

	CheckRecipes(o, t)
}

func TestDelete(t *testing.T) {
	o, err := Delete(recipeId)

	require.Nil(t, err)
	require.NotNil(t, o)

	_, err = GetById(recipeId)

	require.NotNil(t, err)
}

func CheckRecipes(recipes []models.Recipe, t *testing.T) {
	for _, r := range recipes {
		require.Equal(t, recipeId, r.Id)
		require.Equal(t, StrValues["title"], r.Title)
		require.Equal(t, IntValues["time"], r.Time)
		require.Equal(t, IntValues["people"], r.People)
		require.Equal(t, StrValues["ingredients"], r.Ingredients)
		require.Equal(t, StrValues["instructions"], r.Instructions)
	}
}
