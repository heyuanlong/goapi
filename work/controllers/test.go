package controllers

import (
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("goapi!\n"))
}

func Test1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("goapi test1!\n"))
}
