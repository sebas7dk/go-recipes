package config

import "github.com/joho/godotenv"

//Env environemnt
var Env map[string]string

//NewConfig read the .env file
func NewConfig(path string) error {
	var err error

	Env, err = godotenv.Read(path)

	return err
}

// Get the config variable
func Get(field string) string {
	return Env[field]
}
