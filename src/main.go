package main

import (
	"fmt"
	"log"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/users"
)

func cors(w http.ResponseWriter, r *http.Request) {
	//设置监听的端口
	fmt.Println("hhh")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	r.ParseForm()
	fmt.Println("Recieve request: ", r.Form)
	//这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	//设置访问的路由
	http.HandleFunc("/", cors)
	http.HandleFunc("/users/login", users.Register)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
