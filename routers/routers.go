package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
