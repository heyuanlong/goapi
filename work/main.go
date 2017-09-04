package main

import (
	"log"
	"net/http"
	_ "goapi/work/initial"

	krouter "goapi/work/routers"

)

func main() {
	log.Fatal(http.ListenAndServe(":8000", krouter.R))
}