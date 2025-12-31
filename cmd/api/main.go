package main

import (
	"log"

	"github.com/devlongs/collective/internal/env"
	"github.com/devlongs/collective/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		storage: store,
		config:  cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
