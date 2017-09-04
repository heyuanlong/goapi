package controllers

import (
	"net/http"
	"goapi/lib/utils"
	klog "goapi/lib/log"
)

func Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("goapi!\n"))
}

func Test1(w http.ResponseWriter, r *http.Request) {

	klog.Klog.Println(utils.GetTimes())
	klog.Klog.Println(utils.GetTimesString())
	klog.Klog.Println(utils.GetTimesNano())
	w.Write([]byte("goapi test1!\n"))
}

func Testjson(w http.ResponseWriter, r *http.Request) {
	str := utils.GetErrResponse(201,"fail")
	w.Write([]byte(str))
}