package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/morzhanov/go-elk-example/internal/config"
)

const initErr = "initialization error in step %s: %w"

func failOnError(step string, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf(initErr, step, err))
	}
}

func main() {
	c, e := config.NewConfig()
	failOnError("config", e)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	log.Println("App successfully started!")
	<-quit
	log.Println("received os.Interrupt, exiting...")
}
