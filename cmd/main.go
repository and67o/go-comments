package main

import (
	"context"
	"flag"
	"github.com/and67o/go-comments/internal/app"
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/and67o/go-comments/internal/server/internalhttp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	config, err := configuration.New(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var ch = make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(context.Background(), config.Server.Timeout.Server)
	defer cancel()

	a := app.New(config)

	httpServer := internalhttp.New(a)

	watchSignal(ch, httpServer, ctx)
}

func watchSignal(ch chan os.Signal, httpServer interfaces.HTTPApp, ctx context.Context )  {
	signal.Notify(ch, os.Interrupt, syscall.SIGTSTP)

	log.Println("server start")

	go func() {
		err := httpServer.Start()
		if err != nil {
			log.Fatalf("Server failed to start err: %v", err)
		}
	}()

	interrupt := <-ch

	log.Printf("Server is shutting down due to %+v\n", interrupt)

	err := httpServer.Stop(ctx)
	if err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
