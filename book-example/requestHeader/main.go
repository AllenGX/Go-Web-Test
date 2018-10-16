package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header                   // header is a map
	accept := h["Accept-Encoding"]  //[gzip, deflate, br]
	acc := h.Get("Accept-Encoding") //	gzip, deflate, br
	fmt.Fprintln(w, accept)
	fmt.Fprintln(w, acc)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}
