package config

import (
	"database/sql"
	"fmt"
)

func FileTest() bool {
	fmt.Println("[TEST]: Config File Checking...")
	_, err := C.GetConfig()
	if err != nil {
		fmt.Println("[FAIL]:" + err.Error())
		return false
	} else {
		fmt.Println("[PASS]: Config File Available and Accepted.")
		return true
	}
}

func MySQLTest() bool {
	fmt.Println("[TEST]: MySQL Database Checking...")
	_, err := sql.Open("mysql", C.Authenticators.SQLUserName + ":" + C.Authenticators.SQLPassword + "@tcp("+C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Println("[FAIL]: " + err.Error())
		return false
	} else {
		fmt.Println("[PASS]: MySQL is ready to use.")
		return true
	}
}
