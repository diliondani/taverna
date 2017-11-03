package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/diliondani/taverna/tavernawebportal"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address ", config.ServerAddress)
	tavernawebportal.RunWebPortal(config.ServerAddress)
}
