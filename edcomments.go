package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BruzGus/edcomment/migration"
	"github.com/BruzGus/edcomment/routes"
	"github.com/urfave/negroni"
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
	//inicializa las rutas del proyecto
	router := routes.InitRoutes()

	// Inicia los  middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    "8585",
		Handler: n,
	}
	log.Println("iniciado el servidor en http://localhost:8585")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}
