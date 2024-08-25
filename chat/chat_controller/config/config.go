package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type Config struct {
	DB struct {
		Database string
		URL      string
	}

	Kafka struct {
		URL      string
		ClientID string
	}

	Server struct {
		Port string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		log.Fatalf("failed init config, err : %v", err)
	} else if err = toml.NewDecoder(f).Decode(c); err != nil {
		log.Fatalf("failed inject config data, err : %v", err)
	}

	return c
}
