package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"model"
	"net/http"
	"strconv"
)

type ItemsRet struct {
	ItemCnt int          `json:"ItemsCount"`
	Items   []model.Item `json:"Items"`
}

type ItemRet struct {
	ItemName    string `json:"ItemName"`
	Category    string `json:"Category"`
	Price       int    `json:"Price"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
	SellerID    int    `json:"SellerID"`
}

type CreateItemRet struct {
	ItemID int `json:"ItemID"`
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseMultipartForm(32 << 20)
	cookie, err := r.Cookie("SmartChainToken")
	if err != nil {
		if err == http.ErrNoCookie {
			info := CreateItemRet{-1}
			ret, _ := json.Marshal(info)
			fmt.Fprint(w, string(ret))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	//enc := mahonia.NewEncoder("GBK")
	sellerid := CheckSession(cookie.Value)
	itemname := r.MultipartForm.Value["ItemName"][0]

	fmt.Println(itemname)
	
	category := r.MultipartForm.Value["Category"][0]
	fmt.Println(category)
	price, _ := strconv.Atoi(r.MultipartForm.Value["Price"][0])
	description := r.MultipartForm.Value["Description"][0]
	image := r.MultipartForm.Value["Image"][0]
	item := model.Item{0,itemname,description,price,category,image,sellerid}
	itemid := model.CreateNewItem(item)
	info := CreateItemRet{itemid}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
}

func ItemsRetrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	var name string
	var priceUb int
	var priceLb int
	var category string
	if len(r.Form["name"]) > 0 {
		name = r.Form["name"][0]
	} else {
		name = "%"
	}
	if len(r.Form["price_ub"]) > 0 {
		priceUb, _ = strconv.Atoi(r.Form["price_ub"][0])
	} else {
		priceUb = 1 << 32 -1
	}
	if len(r.Form["price_lb"]) > 0 {
		priceLb, _ = strconv.Atoi(r.Form["price_lb"][0])
	} else {
		priceLb = -1
	}
	if len(r.Form["category"]) > 0 {
		category = r.Form["category"][0]
	} else {
		category = "%"
	}
	items, _ := model.GetItems( "%" + name + "%",priceLb,priceUb,"%" + category + "%")
	num := len(items)
	info := ItemsRet{num,items}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
}

func GetItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()
	id, _ := strconv.Atoi(mux.Vars(r)["ItemID"])
	item, _ := model.GetItem(id)
	info := ItemRet{item.ItemName,item.Category,item.Price,item.Description,item.Image,item.SellerID}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
}