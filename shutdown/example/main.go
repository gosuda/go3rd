package main

import (
	"context"
	"log"
	"os"
	"time"

	"gosuda.org/go3rd/shutdown"
)

func main() {
	logger := log.New(os.Stdout, "app: ", log.LstdFlags)

	// Set up signal handling
	ctx := shutdown.HandleKillSig(context.Background(), func() {
		logger.Println("Performing cleanup tasks...")
		time.Sleep(2 * time.Second) // example cleanup work
		logger.Println("Cleanup completed.")
		logger.Println("Application has exited gracefully.")
	}, logger)

	logger.Println("Application is running. Press Ctrl+C to exit.")
	<-ctx.Done()

}
