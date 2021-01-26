package config

import (
	"fmt"
	"os"
)

const (
	KEY_ENV_APIKEY string = "TWAPICLIENT_APIKEY"
	KEY_ENV_HOST   string = "TWAPICLIENT_HOST"
)

type Config struct {
	APIKey string
	Host   string
}

func LoadConf() (*Config, error) {
	var cnf Config
	cnf.APIKey = os.Getenv(KEY_ENV_APIKEY)
	if len(cnf.APIKey) == 0 {
		return nil, fmt.Errorf("Missing api key in environment")
	}
	cnf.Host = os.Getenv(KEY_ENV_HOST)
	if len(cnf.Host) == 0 {
		return nil, fmt.Errorf("Missing base url")
	}
	return &cnf, nil
}
