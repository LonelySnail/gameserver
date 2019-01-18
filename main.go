package main

import (
	"gameserver/bin/conf"
	"gameserver/router"
	"net/http"
	"time"

)

func main() {
	rout := router.InitRouter()
	s := &http.Server{
		Addr:           conf.Config.Host,
		Handler:        rout,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}