package model

import (
	"config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Individual struct {
	User
	Email string
	Tel   int
}

func NewIndividual(individual Individual) string {
	config.C.GetConfig()
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return "Failed to open database:" + err.Error()
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO individual(UserName,Password,Email,Tel) VALUES(?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		fmt.Print(err.Error())
		return "Failed to insert:" + err.Error()
	}
	res, err := stmt.Exec(individual.UserName, individual.Password, individual.Email, individual.Tel)
	if err != nil {
		fmt.Print(err.Error())
		return "There's problem in input data set:" + err.Error()
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Print(err.Error())
		return "ID is not correct:" + err.Error()
	}

	query, _ := db.Prepare("SELECT UserID FROM individual WHERE UserName = ?")
	rows, err := query.Query(individual.UserName)
	defer query.Close()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(id)
	}
	return strconv.FormatInt(id, 10)
}
