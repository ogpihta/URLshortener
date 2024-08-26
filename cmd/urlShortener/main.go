package main

import (
	"URLshortener/internal/parserConfig"
	"log"
	"os"
)

func main() {

	configPath := "../../config/local.yaml"
	err := os.Setenv("CONFIG_PATH", configPath)
	if err != nil {
		log.Fatalf("Failed to set CONFIG_PATH: %v", err)
	}
	cfg := parserConfig.MustLoad()

}
