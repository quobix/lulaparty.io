package routers

import (
	//"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	//"github.com/quobix/lulaparty.io/security"
	//"github.com/quobix/lulaparty.io/controllers"
	"net/http"
	"github.com/quobix/lulaparty.io/model"
)

func SetHelloRoutes(router *mux.Router, ac *model.AppConfig) *mux.Router {
	router.Handle("/",
		//negroni.New(
		//	negroni.HandlerFunc(controllers.HelloController),
		//)).Methods("GET")

		http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	return router
}
