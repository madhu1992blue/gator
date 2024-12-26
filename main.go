package main

import (
	"fmt"
	"log"
	"encoding/json"
	"github.com/madhu1992blue/gator/internal/config"
)
func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Couldn't read config: %v", err)
	}
	
	err = cfg.SetUser("madhusudan")
	if err != nil {
		log.Fatalf("Couldn't set user: %v", err)
	}
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Couldn't read updated config: %v", err)
	}
	
	dataBytes, err := json.Marshal(updatedCfg)
	
	if err != nil {
		log.Fatalf("Couldn't convert Config to json: %v", err)
	}
	fmt.Println(string(dataBytes))

}
