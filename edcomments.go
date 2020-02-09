package main

import (
	"flag"
	"log"

	"github.com/BruzGus/edcomment/migration"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion a la base de datos")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migracion")
		migration.Migrate()
		log.Println("FINALIZO la migracion")
	}
}
