package main

import "log"

func main() {
	cfg := config{
		addr: ":8000",
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Printf("Defined Port %s", cfg.addr)
	log.Fatal(app.run(mux))
}
