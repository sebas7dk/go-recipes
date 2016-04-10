package config

import "github.com/joho/godotenv"

var ENV map[string]string

func NewConfig(path string) error {
	var err error

	ENV, err = godotenv.Read(path)

	return err
}
