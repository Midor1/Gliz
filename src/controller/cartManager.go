package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"strconv"
)

type CartInfoRet struct {
	CartStatus int `json:"status"`
}

type CartRet struct {
	ItemsCount int          `json:"ItemsCount"`
	CartItems  []model.CartItem `json:"CartItems"`
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseMultipartForm(32 << 20)
	cookie, err := r.Cookie("SmartChainToken")
	if err != nil {
		if err == http.ErrNoCookie {
			info := CartInfoRet{-1}
			ret, _ := json.Marshal(info)
			fmt.Fprint(w, string(ret))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	activeid := CheckSession(cookie.Value)
	itemid, _ := strconv.Atoi(r.MultipartForm.Value["ItemID"][0])
	var amount int
	if len(r.MultipartForm.Value["Amount"]) > 0 {
		amount, _ = strconv.Atoi(r.MultipartForm.Value["Amount"][0])
	} else {
		amount = 1
	}
	sat := model.AddToCart(activeid,itemid,amount)
	if sat {
		info := CartInfoRet{0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w,string(ret))
	} else {
		info := CartInfoRet{-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w,string(ret))
	}

}

func CartRetrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseMultipartForm(32 << 20)
	cookie, err := r.Cookie("SmartChainToken")
	if err != nil {
		if err == http.ErrNoCookie {
			info := CartInfoRet{-1}
			ret, _ := json.Marshal(info)
			fmt.Fprint(w, string(ret))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	activeid := CheckSession(cookie.Value)
	items, _ := model.ShowCart(activeid)
	cnt := len(items)
	info := CartRet{cnt,items}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
}

func DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseMultipartForm(32 << 20)
	cookie, err := r.Cookie("SmartChainToken")
	if err != nil {
		if err == http.ErrNoCookie {
			info := CartInfoRet{-1}
			ret, _ := json.Marshal(info)
			fmt.Fprint(w, string(ret))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	activeid := CheckSession(cookie.Value)
	itemid, _ := strconv.Atoi(r.MultipartForm.Value["ItemID"][0])
	var amount int
	if len(r.MultipartForm.Value["Amount"]) > 0 {
		amount, _ = strconv.Atoi(r.MultipartForm.Value["Amount"][0])
	} else {
		amount = 1
	}
	sat := model.DeleteFromCart(activeid,itemid,amount)
	if sat {
		info := CartInfoRet{0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w,string(ret))
	} else {
		info := CartInfoRet{-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w,string(ret))
	}
}