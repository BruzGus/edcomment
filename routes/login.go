package routes

import (
	"github.com/BruzGus/edcomment/controllers"
	"github.com/gorilla/mux"
)

//SetUserRouter ..., router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
