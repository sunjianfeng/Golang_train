package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)

	http.HandleFunc("/login", login)

	err := http.ListenAndServe("localhost:8081", nil)

	if err != nil {
		fmt.Println("ListanAndServer error: ", err.Error())
	}
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	url := req.URL.Path
	if url == "/" {
		w.Write([]byte("hello, world!.\n"))
	} else {

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("pssword:", r.Form["password"])
	}
}
