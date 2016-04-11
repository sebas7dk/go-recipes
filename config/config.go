package config

import "github.com/joho/godotenv"

// ENV environemnt
var ENV map[string]string

//NewConfig ...
func NewConfig(path string) error {
	var err error

	ENV, err = godotenv.Read(path)

	return err
}

// Get returns ..
func Get(field string) (string, bool) {
	data, ok := ENV[field]
	return data, ok
}
