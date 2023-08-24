package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ZookeeperInfo struct {
		Host     string `envconfig:"TEST_ZOOKEEPER_HOST" default:"172.20.0.1:2185"`
		RootNode string `envconfig:"TEST_ZOOKEEPER_ROOT_NODE" default:"/root"`
	}
}

func ReadConfig() (*Config, error) {
	c := new(Config)

	err := envconfig.Process("test", c)
	if err != nil {
		log.Println("[ConfInitialize] failed read config :", err)
		return nil, err
	}

	return c, nil
}
