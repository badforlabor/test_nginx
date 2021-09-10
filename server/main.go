/**
 * Auth :   liubo
 * Date :   2021/9/10 15:28
 * Comment:
 */

package main

import (
	"flag"
	"net/http"
)

var tag = flag.String("tag", "s1", "s1")
var addr = flag.String("addr", ":20210", ":20210")


func main() {
	flag.Parse()

	http.HandleFunc("/api/s1/echo", func(writer http.ResponseWriter, request *http.Request) {
		var a = *tag +" resp:" + request.FormValue("msg")
		writer.Write([]byte(a))
	})
	http.HandleFunc("/api/s2/echo", func(writer http.ResponseWriter, request *http.Request) {
		var a = *tag +" resp:" + request.FormValue("msg")
		writer.Write([]byte(a))
	})
	http.HandleFunc("/api/s3/echo", func(writer http.ResponseWriter, request *http.Request) {
		var a = *tag +" resp:" +  request.FormValue("msg")
		writer.Write([]byte(a))
	})

	http.ListenAndServe(*addr, nil)
}