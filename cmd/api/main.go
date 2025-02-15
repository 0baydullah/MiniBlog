package main

import (
	"log"

	"github.com/0baydullah/MiniBlog/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Printf("Defined Port %s", cfg.addr)
	log.Fatal(app.run(mux))
}
