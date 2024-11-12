package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func HandleKillSig(ctx context.Context, handler func(), logger *log.Logger) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)

	go func() {
		defer func() {
			signal.Stop(sigChannel)
			close(sigChannel)
			cancel()
		}()

		sig := <-sigChannel
		logger.Printf("Received signal %s, shutting down...", sig)
		handler()
	}()

	return ctx
}
