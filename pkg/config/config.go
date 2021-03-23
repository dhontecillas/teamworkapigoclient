package config

import (
	"fmt"
	"net/url"
	"os"
)

const (
	KEY_ENV_APIKEY string = "TWAPICLIENT_APIKEY"
	KEY_ENV_URL    string = "TWAPICLIENT_URL"
)

type Config struct {
	APIKey string
	Scheme string
	Host   string
}

func LoadConf() (*Config, error) {
	var cnf Config
	cnf.APIKey = os.Getenv(KEY_ENV_APIKEY)
	if len(cnf.APIKey) == 0 {
		return nil, fmt.Errorf("Missing api key in environment")
	}
	hostUrl, err := url.Parse(os.Getenv(KEY_ENV_URL))
	if err != nil {
		return nil, err
	}
	cnf.Scheme = hostUrl.Scheme
	cnf.Host = hostUrl.Host
	if len(cnf.Host) == 0 {
		return nil, fmt.Errorf("Missing base url")
	}
	return &cnf, nil
}
