package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<title>Go Web Programming</title>
	</head>
	<body>
	{{ . }}
	</body>
	</html>
	`
	t := template.New("sss")
	t.Parse(html)
	fmt.Println(t.Name())

	//等效
	t, _ = template.ParseFiles("index.html")

	//错误处理办法
	//报错则返回上一层并携带错误信息
	t = template.Must(template.ParseFiles("index.html"))

	fmt.Println(t.Name())
	t.Execute(w, "Hello World!")

	//指定模板
	t.ExecuteTemplate(w, "index.html", "Hello World!")
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
