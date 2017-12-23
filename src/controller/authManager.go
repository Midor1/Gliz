package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"strconv"
)

type AuthRet struct {
	Status int `json:"status"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	r.ParseMultipartForm(32 << 20)
	username := r.MultipartForm.Value["UserName"][0]
	password := r.MultipartForm.Value["Password"][0]
	//Password is hereby plain.
	user := model.User{UserID: 0, UserName: username, Password: password}
	result := model.AuthUser(user)
	id, err := strconv.Atoi(result)
	if err != nil || id == -1 {
		info := AuthRet{-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	} else {
		token := NewSession(id)
		//Default Expiration Time is a week.
		cookie := http.Cookie{Name: "SmartChainToken", Value: token, Path: "/", MaxAge: 86400 * 7}
		http.SetCookie(w, &cookie)
		info := AuthRet{0}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(string(ret))
	}

}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	cookie, err := r.Cookie("SmartChainToken")
	if err != nil {
		info := AuthRet{-1}
		ret, _ := json.Marshal(info)
		fmt.Fprint(w, string(ret))
		fmt.Println(err.Error())
		return
	}
	id := DropSession(cookie.Value)
	rm := http.Cookie{Name: "SmartChainToken", Path: "/", MaxAge: -1}
	http.SetCookie(w, &rm)
	info := AuthRet{id}
	ret, _ := json.Marshal(info)
	fmt.Fprint(w, string(ret))
	fmt.Println(string(ret))
}
