package main

import (
	"log"

	"github.com/ozlemugur/go-clean-arch-tt/config"
	"github.com/ozlemugur/go-clean-arch-tt/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
