package main

import (
	"flag"
	"github.com/and67o/go-comments/internal/app"
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/and67o/go-comments/internal/server/internalhttp"
	"log"
	"os"
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

	a := app.New(&config)

	httpServer := internalhttp.New(a)

	err = httpServer.Start()
	if err != nil {
		//logg.Error("failed to start http server: " + err.Error())
		os.Exit(1)
	}

}
