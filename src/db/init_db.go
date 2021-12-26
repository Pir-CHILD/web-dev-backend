package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:zhao_q777@163.com@tcp(127.0.0.1:3306)/web_dev?multiStatements=true&parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr.Error())
	}
	fmt.Println("Connected!")

	fmt.Println(os.Getwd())
	query, err := ioutil.ReadFile("./db/init_HELPME.sql")
	// fmt.Println(string(query))
	if err != nil {
		fmt.Println(err)
	}
	if _, err := Db.Exec(string(query)); err != nil {
		fmt.Println(0, err)
	}

	query, err = ioutil.ReadFile("./db/init_HELPYOU.sql")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := Db.Exec(string(query)); err != nil {
		fmt.Println(1, err)
	}

	query, err = ioutil.ReadFile("./db/init_USERINFO.sql")
	if err != nil {
		fmt.Println(err)
	}
	_, err = Db.Exec(string(query))
	if err != nil {
		fmt.Println(2, err)
	}

	query, err = ioutil.ReadFile("./db/init_HELPMEDONE.sql")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := Db.Exec(string(query)); err != nil {
		fmt.Println(3, err)
	}

	query, err = ioutil.ReadFile("./db/init_INTERMEDIARYINCOME.sql")
	if err != nil {
		fmt.Println(err)
	}
	if _, err := Db.Exec(string(query)); err != nil {
		fmt.Println(4, err)
	}

	fmt.Println("DB init Done.")
}
