package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/users"
)

func SearchHelpYou(w http.ResponseWriter, r *http.Request) {
	var helpYouInfos []users.HelpYouInfo

	rows, err := db.Db.Query("SELECT * FROM HELPYOU")
	if err != nil {
		fmt.Println("Admin Search HELPYOU error.", err)
	}
	defer rows.Close()

	for rows.Next() {
		var helpYouInfo users.HelpYouInfo
		if err := rows.Scan(&helpYouInfo.HelpYouId, &helpYouInfo.HelpMeId, &helpYouInfo.HelpYouUserId,
			&helpYouInfo.HelpYouDescription, &helpYouInfo.HelpYouCreateTime, &helpYouInfo.HelpYouChangeTime,
			&helpYouInfo.HelpYouState); err != nil {
			fmt.Println("Admin HelpYou Search for row.next() error.", err)
		}
		helpYouInfos = append(helpYouInfos, helpYouInfo)
	}

	fmt.Println("All HelpYouInfos: ", helpYouInfos)
	res, err := json.Marshal(helpYouInfos)
	if err != nil {
		fmt.Println("All HelpYouInfos Marshal error", err)
	}
	w.Write(res)
}
