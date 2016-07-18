package routers

import (
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/quobix/lulaparty.io/security"
	"github.com/quobix/lulaparty.io/controllers"
	"github.com/quobix/lulaparty.io/model"
)

func SetAuthenticationRoutes(router *mux.Router, ac *model.AppConfig) *mux.Router {
	router.HandleFunc("/token-exchange", controllers.Authenticate).Methods("POST")
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(security.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(security.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")
	return router
}
