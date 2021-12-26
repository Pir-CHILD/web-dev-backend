package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/users"
)

func SearchUserInfo(w http.ResponseWriter, r *http.Request) {
	var userInfos []users.UserInfo

	rows, err := db.Db.Query("SELECT * FROM USERINFO")
	if err != nil {
		fmt.Println("Admin Search UserInfo error.", err)
	}

	defer rows.Close()

	for rows.Next() {
		var userInfo users.UserInfo
		if err := rows.Scan(&userInfo.UserId, &userInfo.UserName,
			&userInfo.UserPwd, &userInfo.UserType, &userInfo.UserXingming, &userInfo.CertiType,
			&userInfo.CertiNumber, &userInfo.UserPhone, &userInfo.UserLevel, &userInfo.UserIntroduction,
			&userInfo.UserCity, &userInfo.UserCommunity, &userInfo.RegisterTime, &userInfo.ChangeTime); err != nil {
			fmt.Println("Admin Search UserInfo Select error: ", err)
		}
		userInfos = append(userInfos, userInfo)
	}

	fmt.Println("All UserInfos: ", userInfos)
	res, err := json.Marshal(userInfos)
	if err != nil {
		fmt.Println("All UserInfos Marshal error", err)
	}
	w.Write(res)
}
