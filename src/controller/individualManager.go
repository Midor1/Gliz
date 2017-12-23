package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"strconv"
)

type RegRet struct {
	Status int `json:"status"`
	ID     int `json:"id"`
}

func IndividualRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// New Individual User Register
	r.ParseMultipartForm(32 << 20)
	username := r.MultipartForm.Value["UserName"][0]
	password := r.MultipartForm.Value["Password"][0]
	email := r.MultipartForm.Value["Email"][0]
	tel, err := strconv.Atoi(r.MultipartForm.Value["Tel"][0])
	if err != nil {
		panic(err)
	}
	var individual = model.Individual{model.User{0, username, password}, email, tel}
	var result = model.NewIndividual(individual)
	// If and only if the registration succeeded, the result would be legal integer.
	id, err := strconv.Atoi(result)
	if err != nil {
		info := RegRet{-1, 0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	} else {
		info := RegRet{0, id}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	}

}

func ActiveUserRetrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	cookie, err := r.Cookie("SmartChainToken")

	if err != nil {
		if err == http.ErrNoCookie {
			info := RegRet{-1, 0}
			ret, _ := json.Marshal(info)
			fmt.Fprint(w, string(ret))
			fmt.Println(string(ret))
		} else {
			info := RegRet{-2, 0}
			ret, _ := json.Marshal(info)
			fmt.Fprint(w, string(ret))
			fmt.Println(string(ret))
		}
		return
	}
	userID := CheckSession(cookie.Value)
	if userID == -1 {
		info := RegRet{-3, 0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	} else {
		info := RegRet{0, userID}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	}
}
