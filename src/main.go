package main

import (
	"log"
	"net/http"

	_ "zhaoqing.com/web-dev-backend/src/db"
	_ "zhaoqing.com/web-dev-backend/src/scripts"
	"zhaoqing.com/web-dev-backend/src/users"
)

// 跨域
func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}

func main() {
	//设置访问的路由
	/* 登录 - 注册 */
	http.HandleFunc("/users/register", cors(users.Register))
	http.HandleFunc("/users/login", cors(users.Login))
	/* 劳您驾 */
	http.HandleFunc("/users/helpme/release", cors(users.HelpMeRelease))
	http.HandleFunc("/users/helpme/search", cors(users.HelpMeSearch))
	/* 我可以 */
	http.HandleFunc("/users/helpyou/search", cors(users.HelpYouSearch))
	http.HandleFunc("/users/helpyou/release", cors(users.HelpYouRelease))
	/* 接受 - 删除 */
	http.HandleFunc("/users/helpmedone/delete", cors(users.HelpMeDoneDelete))
	http.HandleFunc("/users/helpmedone/handle", cors(users.HelpMeDoneHandle))
	http.HandleFunc("/users/helpmedone/accept", cors(users.HelpMeDoneAccept))
	// 运行
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
