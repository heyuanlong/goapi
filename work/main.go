package main

import (
	"log"
	"net/http"
	_ "goapi/work/initial"

	krouter "goapi/work/routers"

)

func main() {
	//log.Fatal(http.ListenAndServe(":8000", krouter.R))
	log.Fatal(http.ListenAndServeTLS(":8000","tls/server.crt",
		"tls/server.key", krouter.R))
}