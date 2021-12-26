package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/users"
)

func SearchHelpMe(w http.ResponseWriter, r *http.Request) {
	var helpMeInfos []users.HelpMeInfo

	rows, err := db.Db.Query("SELECT * FROM HELPME")
	if err != nil {
		fmt.Println("Admin Search HELPME error.", err)
	}
	defer rows.Close()

	for rows.Next() {
		var helpMeInfo users.HelpMeInfo
		if err := rows.Scan(&helpMeInfo.HelpMeId, &helpMeInfo.HelpMeUserId, &helpMeInfo.HelpMeType,
			&helpMeInfo.HelpMeTheme, &helpMeInfo.HelpMeDescription, &helpMeInfo.HelpMePeople,
			&helpMeInfo.HelpMeFinishTime, &helpMeInfo.HelpMeIntroPhoto, &helpMeInfo.HelpMeCreateTime,
			&helpMeInfo.HelpMeChangeTime, &helpMeInfo.HelpMeState); err != nil {
			fmt.Println("Admin HelpMe Search for row.next() error.", err)
		}
		helpMeInfos = append(helpMeInfos, helpMeInfo)
	}

	fmt.Println("All HelpMeInfos: ", helpMeInfos)
	res, err := json.Marshal(helpMeInfos)
	if err != nil {
		fmt.Println("All HelpMeInfos Marshal error", err)
	}
	w.Write(res)
}
