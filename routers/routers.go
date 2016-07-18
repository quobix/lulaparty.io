package routers

import (
	"github.com/gorilla/mux"
	"github.com/quobix/lulaparty.io/model"
)

func InitRoutes(ac *model.AppConfig) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = SetHelloRoutes(router, ac)
	router = SetAuthenticationRoutes(router, ac)
	return router
}
