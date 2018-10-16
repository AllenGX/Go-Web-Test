package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// 用来得到file 文件 和  其他键值对数据
	r.ParseMultipartForm(1024)                //设置需要提取数据的大小
	fmt.Fprintln(w, r.MultipartForm)          //提交方式必须是multipart/form-data
	fmt.Fprintln(w, r.FormValue("hello"))     // 得到get数据
	fmt.Fprintln(w, r.PostFormValue("hello")) // 得到Post数据
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
