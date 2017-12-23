package model

import (
	"config"
	"database/sql"
	"fmt"
	"strconv"
)

type User struct {
	UserID   int
	UserName string
	Password string
}

func AuthUser(user User) string {
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return "Failed to open database:" + err.Error()
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT UserID FROM individual WHERE UserName = ? AND Password = ?")
	defer stmt.Close()
	if err != nil {
		fmt.Print(err.Error())
		return "Failed to insert:" + err.Error()
	}
	rows, err := stmt.Query(user.UserName, user.Password)
	if err != nil {
		fmt.Print(err.Error())
		return "There's problem in input data set:" + err.Error()
	}
	UserID := -1
	for rows.Next() {
		err := rows.Scan(&UserID)
		if err != nil {
			return "SQL Internal error:" + err.Error()
		}
	}
	return strconv.Itoa(UserID)
}
