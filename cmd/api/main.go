package main

import (
	"log"

	"github.com/devlongs/collective/internal/db"
	"github.com/devlongs/collective/internal/env"
	"github.com/devlongs/collective/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn:          env.GetString("DB_DSN", "postgres://admin:adminpassword@localhost/collective?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "10m"),
		},
	}

	db, err := db.New(cfg.db.dsn, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("DB connected succeessfully")

	store := store.NewStorage(db)

	app := &application{
		storage: store,
		config:  cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
