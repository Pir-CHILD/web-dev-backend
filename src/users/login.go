package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aaa")
	var params map[string]string // 存储参数
	err := json.NewDecoder(r.Body).Decode(&params)
	fmt.Println(params)
	if err != nil {
		fmt.Println(err)
	}

}
