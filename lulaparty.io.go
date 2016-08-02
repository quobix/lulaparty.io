package main

import (
	"github.com/quobix/lulaparty.io/routers"
	"net/http"
	"github.com/codegangsta/negroni"

	"github.com/goinggo/tracelog"
	"github.com/quobix/lulaparty.io/data"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	ac := data.CreateAppConfig(true)


	router := routers.InitRoutes(ac)
	n := negroni.Classic()
	n.UseHandler(router)
	tracelog.Trace("main","main","Starting Lulu Service")
	http.ListenAndServe(":5000", n)

}
