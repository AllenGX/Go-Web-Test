package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	// writer.WriteHeader(200)
	// a := []byte{'1', '2'}
	// b := []string{"111", "222"}
	// writer.Write(a)
	// header := writer.Header()
	// // header["a"] = b
	// header.Add("c", "ccc")
	// fmt.Println()

	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
