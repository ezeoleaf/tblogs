package main

import (
	"log"

	"github.com/ezeoleaf/tblogs/internal/app"
	"github.com/ezeoleaf/tblogs/internal/config"
)

func main() {
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	a := app.NewApp(cfg)
	a.Run()
}
