package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/scripts"
)

func HelpYouSearch(w http.ResponseWriter, r *http.Request) {
	var helpMeInfos []HelpMeInfo
	var params map[string]string

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println("HelpMeSearch params decode error.", err)
	}
	fmt.Println("HelpMeSearch params: ", params)
	// 找出所有非自己发出的请求
	rows, err := db.Db.Query("SELECT * FROM HELPME WHERE helpMeUserId<>?", params["userId"])
	if err != nil {
		fmt.Println("HelpMeSearch error.", err)
	}
	defer rows.Close()

	for rows.Next() {
		var helpMeInfo HelpMeInfo
		if err := rows.Scan(&helpMeInfo.HelpMeId, &helpMeInfo.HelpMeUserId, &helpMeInfo.HelpMeType,
			&helpMeInfo.HelpMeTheme, &helpMeInfo.HelpMeDescription, &helpMeInfo.HelpMePeople,
			&helpMeInfo.HelpMeFinishTime, &helpMeInfo.HelpMeIntroPhoto, &helpMeInfo.HelpMeCreateTime,
			&helpMeInfo.HelpMeChangeTime, &helpMeInfo.HelpMeState); err != nil {
			fmt.Println("HelpMeSearch for row.next() error.", err)
		}
		helpMeInfos = append(helpMeInfos, helpMeInfo)
	}

	fmt.Println("HelpMeSearch info:", helpMeInfos)
	res, err := json.Marshal(helpMeInfos)
	if err != nil {
		fmt.Println("HelpMeSearch Marshal error", err)
	}
	w.Write(res)
}

func HelpYouRelease(w http.ResponseWriter, r *http.Request) {
	var helpYouInfo HelpYouInfo

	err := json.NewDecoder(r.Body).Decode(&helpYouInfo)
	if err != nil {
		fmt.Println("HelpYouRelease decode error.", err)
	}
	fmt.Println("HelpYouRelease ", helpYouInfo)

	if err := db.Db.QueryRow("SELECT * FROM HELPYOU WHERE helpMeId=? and helpYouUserId=?",
		helpYouInfo.HelpMeId, helpYouInfo.HelpYouUserId).Scan(&helpYouInfo.HelpMeId,
		&helpYouInfo.HelpYouUserId); err != nil {
		if err == sql.ErrNoRows { // 可以申请\
			// 插入 helpyou
			tId := scripts.GenId()
			result, err := db.Db.Exec(`INSERT INTO HELPYOU (helpYouId, helpMeId, 
				helpYouUserId, helpYouDescription, helpYouCreateTime, helpYouChangeTime, 
				helpYouState) VALUES (?, ?, ?, ?, ?, ?, ?)`,
				tId, helpYouInfo.HelpMeId, helpYouInfo.HelpYouUserId,
				helpYouInfo.HelpYouDescription, helpYouInfo.HelpYouCreateTime,
				helpYouInfo.HelpYouChangeTime, helpYouInfo.HelpYouState)
			if err != nil {
				fmt.Println("HelpYouRelease INSERT error!", err)
			}
			id, err := result.LastInsertId()
			if err != nil {
				fmt.Println("HelpYouRelease LastInsertId error!", err)
			}
			fmt.Println("HelpYouRelease INSERT LastInsertId: ", id)

			w.WriteHeader(200)
		} else { // 重复的 HelpYou 申请

			fmt.Println("HelpYouRelease failed. Existed release.", err)
			w.WriteHeader(400)
			w.Write([]byte("Existed release."))
		}
	} else {
		fmt.Println("HelpYouRelease failed! something error.", err)
		w.WriteHeader(500)
		w.Write([]byte("erooooooooooooooooor."))
	}
}
