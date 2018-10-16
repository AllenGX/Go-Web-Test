package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()                        //解析参数，默认不解析
	fmt.Println(req.Form)                  //form的信息
	fmt.Println("path:", req.URL.Path)     //path
	fmt.Println("scheme:", req.URL.Scheme) //
	fmt.Println(req.Form["url_long"])

	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "|"))
	}
	fmt.Fprintf(w, "Hello Allen")
	fmt.Println("path:", req.URL.Path)
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
