package main

import (
	"fmt"
	"net/http"
)
/*
"net/http":
	-> handlesFunction(http.ResponseWriter, http.Request)
		-> ResponseWrite is used to fill in the HTTP response
*/
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

//read all HTTP request readers and achoing then into the response body
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	//register the handles on server routes
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}