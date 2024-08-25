package parserConfig

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-required:"true"`
	StoragePath string     `yaml:"storagePath" env-required:"true"`
	GRPC        GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	configPath := os.Getenv(key:"CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}

