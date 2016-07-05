package main

import (
	"github.com/quobix/lulaparty.io/routers"
	"net/http"
	"github.com/codegangsta/negroni"

	"github.com/goinggo/tracelog"
)

func main() {

	tracelog.Start(tracelog.LevelTrace)
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	tracelog.Trace("main","main","Starting Lulu Service")
	http.ListenAndServe(":5000", n)

}
