package routers

import (
	"github.com/gorilla/mux"
	kcontrol "goapi/work/controllers"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()
	addRouter()
}

func addRouter()  {
	R.HandleFunc("/", kcontrol.Test)
	R.HandleFunc("/test1", kcontrol.Test1)
}
