package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/scripts"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var userInfo UserInfo
	fmt.Println("r.Body: ", r.Body)

	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		fmt.Println("Register decode err: ", err)
	}
	fmt.Println("Register params: ", userInfo)

	if err := db.Db.QueryRow("SELECT * FROM USERINFO WHERE userName=?",
		userInfo.UserName).Scan(&userInfo.UserName); err != nil {
		if err == sql.ErrNoRows { // 不存在的数据
			tId := scripts.GenId()
			result, err := db.Db.Exec(`INSERT INTO USERINFO (userId, userName, userPwd, 
				userType, userXingming, certiType, certiNumber, userPhone, userLevel, 
				userIntroduction, userCity, userCommunity, registerTime, changeTime) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				tId, userInfo.UserName, userInfo.UserPwd, userInfo.UserType,
				userInfo.UserXingming, userInfo.CertiType, userInfo.CertiNumber,
				userInfo.UserPhone, userInfo.UserLevel, userInfo.UserIntroduction,
				userInfo.UserCity, userInfo.UserCommunity, userInfo.RegisterTime, userInfo.ChangeTime)
			if err != nil {
				fmt.Println("Register INSERT error!", err)
			}
			id, err := result.LastInsertId()
			if err != nil {
				fmt.Println("Register LastInsertId error!", err)
			}
			fmt.Println("Register INSERT LastInsertId: ", id)
			w.WriteHeader(200)
		} else {
			fmt.Println("Register failed! Existed username.", err)
			w.WriteHeader(400)
			w.Write([]byte("Existed username."))
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userInfo UserInfo
	var params map[string]string // 存储参数
	err := json.NewDecoder(r.Body).Decode(&params)
	fmt.Println(params)
	if err != nil {
		fmt.Println(err)
	}

	if err := db.Db.QueryRow("SELECT * FROM USERINFO WHERE userName=? and userPwd=?",
		params["username"], params["password"]).Scan(&userInfo.UserId, &userInfo.UserName,
		&userInfo.UserPwd, &userInfo.UserType, &userInfo.UserXingming, &userInfo.CertiType,
		&userInfo.CertiNumber, &userInfo.UserPhone, &userInfo.UserLevel, &userInfo.UserIntroduction,
		&userInfo.UserCity, &userInfo.UserCommunity, &userInfo.RegisterTime, &userInfo.ChangeTime); err != nil {
		if err == sql.ErrNoRows { // 用户名或密码错误
			fmt.Println("Login failed. No such username and pwd.")
			w.WriteHeader(400)
			w.Write([]byte("Username or password error."))
		} else { // 应该不会进入这个 else ...
			fmt.Println("Something error.")
			w.WriteHeader(500)
			w.Write([]byte("Username or password error."))
		}
	} else {
		fmt.Println("user info:", userInfo)
		res, err := json.Marshal(userInfo)
		if err != nil {
			fmt.Println("Userinfo Marshal error", err)
		}
		w.Write(res)
	}
}
