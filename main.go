package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"query-language/database"
	"sync"
	"syscall"
)

var (
	waitGroup *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {

	// Get current working directory
	//
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Read contents of configuration file
	//
	configurationPath := cwd + "/configuration.yaml"
	configurationBytes, err := os.ReadFile(configurationPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create database configuration from yaml bytes
	// and create a new database instance
	//
	configuration, err := database.ConfigurationFromYaml(configurationBytes)
	if err != nil {
		log.Fatal(err)
		return
	}
	db := database.New(configuration)

	// Context preparation
	//
	ctx, cancelFunc := context.WithCancel(context.Background())

	// Start execution of database.
	//
	waitGroup.Add(1)
	go db.Run(waitGroup, ctx, cancelFunc)

	// Stop execution on signals
	//
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	cancelFunc()

	// Wait for all routines to finish
	//
	waitGroup.Wait()
}
