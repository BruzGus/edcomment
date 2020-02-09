package routes

import (
	"github.com/BruzGus/edcomment/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetCommentRouter ..., maneja la ruta de los Comentarios
func SetCommentRouter(router *mux.Router) {
	prefix := "/api/comments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
