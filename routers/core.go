package routers

import (
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/quobix/lulaparty.io/security"
	"github.com/quobix/lulaparty.io/controllers"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(security.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}