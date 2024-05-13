package main

import (
	"log"

	"github.com/akshay0074700747/order-inventory_management/config"
	dependencyinjection "github.com/akshay0074700747/order-inventory_management/dependency_injection"
)

func main() {

	//loading the configurations from the env
	cfg, err := config.LoadConfigurationss()
	if err != nil {
		log.Fatalf("fatal error %s , exiting...", err.Error())
	}

	//this application in built in clean architecture
	//initialising all the dependencies
	controller := dependencyinjection.InjectDependencies(cfg)

	//starting up the server in the specified port
	controller.Start(cfg.Port)
}
