package model

import (
	"config"
	"database/sql"
	"fmt"
)

type Cart struct {
	UserID    int
	CartItems []CartItem
}

type CartItem struct {
	ItemID int `json:"ItemID"`
	Amount int `json:"Amount"`
}

func AddToCart(id int,item int,amount int) bool {
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT CartID FROM cart WHERE ItemID = ? AND UserID = ?")
	rows, err := stmt.Query(item,id)
	defer stmt.Close()
	defer rows.Close()
	var eid = -1
	for rows.Next() {
		rows.Scan(&eid)
	}
	if eid == -1 {
		query, err := db.Prepare("INSERT INTO cart(ItemID,UserID,Amount) VALUES(?,?,?)")
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res, err := query.Query(item,id,amount)
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res.Close()
		return true
	} else {
		query, err := db.Prepare("UPDATE cart SET Amount = Amount + ? WHERE ItemID = ? AND UserID = ?")
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res, err := query.Query(amount,item,id)
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res.Close()
		return true
	}
}

func ShowCart(id int) ([]CartItem, error) {
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	defer db.Close()
	query, err := db.Prepare("SELECT ItemID,Amount FROM cart WHERE UserID = ?")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	rows, err := query.Query(id)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	defer query.Close()
	defer rows.Close()
	var items []CartItem
	var item CartItem
	for rows.Next() {
		rows.Scan(&item.ItemID,&item.Amount)
		items = append(items, item)
	}
	return items, nil
}

func DeleteFromCart(id int,item int,amount int) bool {
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT Amount FROM cart WHERE ItemID = ? AND UserID = ?")
	rows, err := stmt.Query(item,id)
	defer stmt.Close()
	defer rows.Close()
	var eid = -1
	for rows.Next() {
		rows.Scan(&eid)
	}
	if eid <= amount {
		query, err := db.Prepare("DELETE FROM cart WHERE ItemID = ? AND UserID = ?")
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res, err := query.Query(item,id)
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res.Close()
		return true
	} else {
		query, err := db.Prepare("UPDATE cart SET Amount = Amount - ? WHERE ItemID = ? AND UserID = ?")
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res, err := query.Query(amount,item,id)
		if err != nil {
			fmt.Print(err.Error())
			return false
		}
		res.Close()
		return true
	}
}
