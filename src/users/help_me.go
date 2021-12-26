package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/scripts"
)

func HelpMeRelease(w http.ResponseWriter, r *http.Request) {
	var helpMeInfo HelpMeInfo

	err := json.NewDecoder(r.Body).Decode(&helpMeInfo)
	if err != nil {
		fmt.Println("HelpMERelease decode err: ", err)
	}
	fmt.Println("helpMeInfo params: ", helpMeInfo)

	tId := scripts.GenId()
	result, err := db.Db.Exec(`INSERT INTO HELPME (helpMeId, helpMeUserId, helpMeType, 
		helpMeTheme, helpMeDescription, helpMePeople, helpMeFinishtime, helpMeIntroPhoto, 
		helpMeCreateTime, helpMeChangeTime, helpMeState) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		tId, helpMeInfo.HelpMeUserId, helpMeInfo.HelpMeType, helpMeInfo.HelpMeTheme,
		helpMeInfo.HelpMeDescription, helpMeInfo.HelpMePeople, helpMeInfo.HelpMeFinishTime,
		helpMeInfo.HelpMeIntroPhoto, helpMeInfo.HelpMeCreateTime, helpMeInfo.HelpMeChangeTime,
		helpMeInfo.HelpMeState)
	if err != nil {
		fmt.Println("HelpMeRelease INSERT error!", err)
		w.WriteHeader(500)
		w.Write([]byte("HelpMeRelease INSERT error!"))
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("HelpMeRelease LastInsertId error!", err)
	}
	fmt.Println("HelpMeRelease INSERT LastInsertId: ", id)
}

func HelpMeSearch(w http.ResponseWriter, r *http.Request) {
	var helpMeInfos []HelpMeInfo
	var params map[string]string

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println("HelpMeSearch params decode error.", err)
	}
	fmt.Println("HelpMeSearch params: ", params)

	rows, err := db.Db.Query("SELECT * FROM HELPME WHERE helpMeUserId=?", params["userId"])
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
