package main

import (
	"log"
	"os"

	"main/api"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/app/config.yaml"
	}

	wrapperHost := os.Getenv("WRAPPER_HOST")
	if wrapperHost == "" {
		wrapperHost = "wrapper"
	}

	addr := os.Getenv("API_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	server := api.NewServer(configPath, wrapperHost)
	log.Fatal(server.Start(addr))
}
