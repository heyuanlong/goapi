package main

import (
	"log"
	"net/http"

	krouter "goapi/work/routers"

)

func main() {
	log.Fatal(http.ListenAndServe(":8000", krouter.R))
}