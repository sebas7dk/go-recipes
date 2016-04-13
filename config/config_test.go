package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var v = map[string]string{
	"port":          "APP_PORT",
	"version":       "APP_VERSION",
	"es_domain":     "ES_DOMAIN",
	"es_port":       "ES_PORT",
	"es_index":      "ES_INDEX",
	"es_test_index": "ES_TEST_INDEX",
}

func TestNewConfig(t *testing.T) {
	path1 := ".env"
	path2 := "../.env"

	err := NewConfig(path1)
	require.NotNil(t, err)

	err = NewConfig(path2)
	require.Nil(t, err)
}

func TestENV(t *testing.T) {
	require.NotEmpty(t, Get(v["port"]))
	require.NotEmpty(t, Get(v["version"]))
	require.NotEmpty(t, Get(v["es_domain"]))
	require.NotEmpty(t, Get(v["es_port"]))
	require.NotEmpty(t, Get(v["es_index"]))
	require.NotEmpty(t, Get(v["es_test_index"]))
}
