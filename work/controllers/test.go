package controllers

import (
	"net/http"
	"goapi/lib/utils"
	klog "goapi/lib/log"

	"fmt"
	//"strings"
)

func Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("ke y:", k)
		lens :=  len(v[0])
		for i:=0; i < lens ;i++  {
			fmt.Println(i,v[0][i])
		}
	}
	//fmt.Fprintf(w, "Hello go web user! \n") //这个写入到w的是输出到客户端的

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