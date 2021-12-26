package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
)

func HelpMeDoneHandle(w http.ResponseWriter, r *http.Request) {
	var helpYouInfos []HelpYouInfo
	var params map[string]string

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println("HelpMeDoneHandle params error: ", err)
	}
	fmt.Println("HelpMeDoneHandle params: ", params)

	// 在 HELPYOU 中找到对应 HELPME
	rows, err := db.Db.Query("SELECT * FROM HELPYOU WHERE helpMeId=?", params["helpMeId"])
	if err != nil {
		fmt.Println("HelpMeDoneHandle error.", err)
	}
	defer rows.Close()

	for rows.Next() {
		var helpYouInfo HelpYouInfo
		if err := rows.Scan(&helpYouInfo.HelpYouId, &helpYouInfo.HelpMeId, &helpYouInfo.HelpYouUserId,
			&helpYouInfo.HelpYouDescription, &helpYouInfo.HelpYouCreateTime, &helpYouInfo.HelpYouChangeTime,
			&helpYouInfo.HelpYouState); err != nil {
			fmt.Println("HelpMeDoneHandle Search for row.next() error.", err)
		}
		helpYouInfos = append(helpYouInfos, helpYouInfo)
	}

	fmt.Println("HelpMeDone Search info:", helpYouInfos)
	res, err := json.Marshal(helpYouInfos)
	if err != nil {
		fmt.Println("HelpMeDone Search Marshal error", err)
	}
	w.Write(res)
}

func HelpMeDoneAccept(w http.ResponseWriter, r *http.Request) {
	var helpMeDone HelpMeDone

	err := json.NewDecoder(r.Body).Decode(&helpMeDone)
	if err != nil {
		fmt.Println("HelpMeDone decode err: ", err)
	}
	fmt.Println("helpMeDone params: ", helpMeDone)

	result, err := db.Db.Exec(`INSERT INTO HELPMEDONE (helpMeId, helpMeUserId, 
		helpYouUserId, helpYouFinishTime, helpMeUserSpend, helpYouUserSpend)
		VALUES (?, ?, ?, ?, ?, ?)`, helpMeDone.HelpMeId, helpMeDone.HelpMeUserId,
		helpMeDone.HelpYouUserId, helpMeDone.HelpYouFinishTime,
		helpMeDone.HelpMeUserSpend, helpMeDone.HelpYouUserSpend)
	if err != nil {
		fmt.Println("HelpMeDone INSERT error!", err)
		w.WriteHeader(500)
		w.Write([]byte("HelpMeDone INSERT error!"))
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("HelpMeDone LastInsertId error!", err)
	}
	fmt.Println("HelpMeDone INSERT LastInsertId: ", id)

	// 更新 helpMeState
	result, err = db.Db.Exec("UPDATE HELPME SET helpMeState=? WHERE helpMeId=?", 0,
		helpMeDone.HelpMeId)
	if err != nil {
		fmt.Println("HelpMeDone HelpMe Update error!", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println("HelpMeDone HelpMeState Update LastInsertId error!", err)
	}
	fmt.Println("HelpMeDone HelpMeState Update LastInsertId: ", id)

	// 更新 helpYouState
	result, err = db.Db.Exec("UPDATE HELPYOU SET helpYouState=? WHERE helpMeId=?", 1,
		helpMeDone.HelpMeId)
	if err != nil {
		fmt.Println("HelpMeDone 111 HelpYou Update error!", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println("HelpMeDone 111 HelpYouState Update LastInsertId error!", err)
	}
	fmt.Println("HelpMeDone 111 HelpYouState Update LastInsertId: ", id)

	result, err = db.Db.Exec("UPDATE HELPYOU SET helpYouState=? WHERE helpMeId<>?", 2,
		helpMeDone.HelpMeId)
	if err != nil {
		fmt.Println("HelpMeDone 222 HelpYou Update error!", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println("HelpMeDone 222 HelpYouState Update LastInsertId error!", err)
	}
	fmt.Println("HelpMeDone 222 HelpYouState Update LastInsertId: ", id)
}

func HelpMeDoneDelete(w http.ResponseWriter, r *http.Request) {
	var params map[string]string

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println("HelpMeDoneDelete params error: ", err)
	}
	fmt.Println("HelpMeDoneDelete params: ", params)

	// 更新 helpMeState
	result, err := db.Db.Exec("UPDATE HELPME SET helpMeState=? WHERE helpMeId=?", 2, params["helpMeId"])
	if err != nil {
		fmt.Println("HelpMe Update error!", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("HelpMeState Update LastInsertId error!", err)
	}
	fmt.Println("HelpMeState Update LastInsertId: ", id)

	// 更新 helpYouState
	result, err = db.Db.Exec("UPDATE HELPYOU SET helpYouState=? WHERE helpMeId=?", 3, params["helpMeId"])
	if err != nil {
		fmt.Println("HelpYou Update error!", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println("HelpYouState Update LastInsertId error!", err)
	}
	fmt.Println("HelpYouState Update LastInsertId: ", id)
}
