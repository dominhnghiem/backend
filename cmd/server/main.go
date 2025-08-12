package main

import (
	"log"

	"your.module/name/bootstrap"
)

func main() {
	cfg := bootstrap.LoadConfig()
	db := bootstrap.NewDB(cfg)

	app := bootstrap.NewAppWithDeps(db, cfg)
	if err := bootstrap.Run(app, cfg); err != nil {
		log.Fatal(err)
	}
}
