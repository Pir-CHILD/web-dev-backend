package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"zhaoqing.com/web-dev-backend/src/db"
	"zhaoqing.com/web-dev-backend/src/users"
)

func SearchIncome(w http.ResponseWriter, r *http.Request) {
	var types [4]int
	var helpMeDones []users.HelpMeDone

	rows, err := db.Db.Query("SELECT * FROM HELPMEDONE")
	if err != nil {
		fmt.Println("Admin Search Income error.", err)
	}
	defer rows.Close()

	for rows.Next() {
		var helpMeDone users.HelpMeDone
		if err := rows.Scan(&helpMeDone.HelpMeId, &helpMeDone.HelpMeUserId,
			&helpMeDone.HelpYouUserId, &helpMeDone.HelpYouFinishTime,
			&helpMeDone.HelpMeUserSpend, &helpMeDone.HelpYouUserSpend); err != nil {
			fmt.Println("Admin Search HelpMeDone Select error: ", err)
		}
		helpMeDones = append(helpMeDones, helpMeDone)
	}

	for _, item := range helpMeDones {
		var helpMeType int

		if err = db.Db.QueryRow("SELECT helpMeType FROM HELPME WHERE helpMeId=?",
			item.HelpMeId).Scan(&helpMeType); err != nil {
			fmt.Println("Admin Search HelpMeType error:", err)
		} else {
			types[helpMeType] += item.HelpMeUserSpend
			types[helpMeType] += item.HelpYouUserSpend
		}
	}

	fmt.Println("All Incomes: ", types)
	res, err := json.Marshal(types)
	if err != nil {
		fmt.Println("All Incomes Marshal error", err)
	}
	w.Write(res)
}
