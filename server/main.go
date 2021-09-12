/**
 * Auth :   liubo
 * Date :   2021/9/10 15:28
 * Comment:
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
)

var tag = flag.String("tag", "s1", "s1")
var addr = flag.String("addr", ":20210", ":20210")


func main() {
	flag.Parse()

	_myHandler.Reg("/api/s1/echo", func(writer http.ResponseWriter, request *http.Request) {
		var a = *tag +" resp:" + request.FormValue("msg")
		writer.Write([]byte(a))
	})
	_myHandler.Reg("/api/s2/echo", func(writer http.ResponseWriter, request *http.Request) {
		var a = *tag +" resp:" + request.FormValue("msg")
		writer.Write([]byte(a))
	})
	_myHandler.Reg("/api/s3/echo", func(writer http.ResponseWriter, request *http.Request) {
		var a = *tag +" resp:" +  request.FormValue("msg")
		writer.Write([]byte(a))
	})
	_myHandler.Reg("/api/s4", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("error. unknown uri:", request.RequestURI)

		writer.WriteHeader(http.StatusNotFound)
	})

	fmt.Println(*tag, "serve:", *addr)
	http.ListenAndServe(*addr, _myHandler)
}

type myHandler struct {
	actions map[string]http.HandlerFunc
}
func (self *myHandler) Reg(pattern string, handler http.HandlerFunc) {
	self.actions[pattern] = handler
}
func (self *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var uri = r.URL.Path
	var v, ok = self.actions[uri]
	if ok {
		fmt.Println("info. process uri:", r.RequestURI)
		v(w, r)
	} else {
		fmt.Println("error. unknown uri:", r.RequestURI)
		w.WriteHeader(http.StatusNotFound)
	}
}
var _myHandler = &myHandler {make(map[string]http.HandlerFunc) }