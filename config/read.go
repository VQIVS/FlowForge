package config

import (
	"encoding/json"
	"os"
)

func ReadConfig() (Config, error) {
	var c Config
	all, err := os.ReadFile("config.json")
	if err != nil {
		return c, err
	}
	return c, json.Unmarshal(all, &c)
}

func MustReadConfig(configPath string) Config {
	c, err := ReadConfig()
	if err != nil {
		panic(err)
	}
	return c
}
