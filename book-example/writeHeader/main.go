package main

import (
	"fmt"
	"net/http"
)

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.Header().Set("aaa", "aaa")
	w.WriteHeader(200)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/header", headerExample)
	server.ListenAndServe()
}
