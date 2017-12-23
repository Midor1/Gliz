package main

import (
	"config"
	"controller"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	var confNow, _ = config.C.GetConfig()
	fmt.Println("Gliz Server v0.3")
	if !config.FileTest() || !config.MySQLTest() {
		fmt.Println("Quitting...")
		return
	}
	fmt.Println("Server running on: " + confNow.Addresses.SrvAddr)
	rtr := mux.NewRouter()
	//Individual User Register and Active User Retrieve
	rtr.HandleFunc("/individual", controller.IndividualRegister).Methods("POST")
	rtr.HandleFunc("/individual", controller.ActiveUserRetrieve).Methods("GET")
	//Login and logout
	rtr.HandleFunc("/auth", controller.UserLogin).Methods("POST")
	rtr.HandleFunc("/auth", controller.UserLogout).Methods("DELETE")
	//Create Item and Item Request
	rtr.HandleFunc("/item",controller.CreateItem).Methods("POST")
	rtr.HandleFunc("/item",controller.ItemsRetrieve).Methods("GET")

	//SayHello function used to check server status.
	rtr.HandleFunc("/hello", SayHello).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(confNow.Addresses.SrvAddr, nil)
}
