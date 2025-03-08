package main

import (
	"context"
	"github.com/denis-rossati/github-notifications/internal"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signCh

		log.Printf("Received signal %v, shutting down", sig)

		cancel()
	}()

	args, err := internal.GetArgs()

	if err != nil {
		log.Fatalf("An error occurred: %v\n", err)
	}

	internal.Listen(ctx, args.Token)
}
