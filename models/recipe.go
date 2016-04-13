package models

type Recipe struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Category     string `json:"category"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
	Time         int    `json:"time"`
	People       int    `json:"people"`
}

type Recipes []Recipe
