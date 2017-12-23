package model

import (
	"config"
	"database/sql"
	"fmt"
)

type Item struct {
	ItemID      int    `json:"ItemID"`
	ItemName    string `json:"ItemName"`
	Description string `json:"Description"`
	Price       int    `json:"Price"`
	Category    string `json:"Category"`
	Image       string `json:"Image"`
	SellerID    int    `json:"SellerID"`
}

func CreateNewItem(item Item) int {
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return -1
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO item(ItemName,Description,Price,Category,Image,SellerID) VALUES(?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		fmt.Print(err.Error())
		return -1
	}
	res, err := stmt.Exec(item.ItemName, item.Description, item.Price, item.Category, item.Image, item.SellerID)
	if err != nil {
		fmt.Print(err.Error())
		return -1
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Print(err.Error())
		return -1
	}

	query, _ := db.Prepare("SELECT ItemID FROM item WHERE ItemName = ?")
	rows, err := query.Query(item.ItemName)
	defer query.Close()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(id)
	}
	return int(id)
}

func GetItems(name string, priceLb int, priceUb int,category string) ([]Item, error) {
	db, err := sql.Open("mysql", config.C.Authenticators.SQLUserName + ":" + config.C.Authenticators.SQLPassword + "@tcp("+config.C.Addresses.SQLAddr+")/smart?charset=utf8")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	defer db.Close()
	query, err := db.Prepare("SELECT ItemID,ItemName,Price,Description,Category,Image,SellerID FROM item WHERE ItemName LIKE ? AND Price > ? AND Price < ? AND Category LIKE ?")
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	rows, err := query.Query(name, priceLb, priceUb,category)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	defer query.Close()
	defer rows.Close()
	var items []Item
	var item Item
	for rows.Next() {
		rows.Scan(&item.ItemID,&item.ItemName,&item.Price,&item.Description,&item.Category,&item.Image,&item.SellerID)
		items = append(items, item)
	}
	return items, nil
}