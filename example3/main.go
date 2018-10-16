package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
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
}

// session



func login(w http.ResponseWriter, req *http.Request) {
	globalSession, _ := NewManager("memory", "gosessionid", 3600)
	sess := globalSession.SessionStart(w, req)
	req.ParseForm()
	fmt.Println("method:", req.Method)
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, sess.Get("username"))
	} else {
		fmt.Println("username:", req.Form["username"])
		fmt.Println("password:", req.Form["password"])
		sess.Set("username", req.Form["username"])
		http.Redirect(w, req, "/", 302)
	}
}

func main() {
	http.HandleFunc("/sayhello", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
