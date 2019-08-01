package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tzilist/m-service/pkg/server"
)

func main() {
	appServer, serverPort := server.Start(nil, nil)

	log.Printf("Starting Application Server on port [%v]", serverPort)

	// start server and handle graceful shutdown
	go func() {
		if err := appServer.Start(serverPort); err != nil {
			log.Println("Shutting down server")
		}
	}()

	// Gracefull shutdown
	// Add any services that need to be disconnected or drained here
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// stop app server and give 30 seconds to finish up any remaining requests
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := appServer.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Println("Goodbye :)")
	os.Exit(0)
}
