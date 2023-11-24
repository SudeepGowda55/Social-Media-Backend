package main

import (
	"net/http"
	"server/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// router.Use()

	router.HandleFunc("/", api.HomePage).Methods("GET")

	router.HandleFunc("/getusers", api.GetUsers).Methods("GET")
	router.HandleFunc("/createuser", api.CreateUser).Methods("POST")
	router.HandleFunc("/getuser", api.GetUser).Methods("GET")

	router.HandleFunc("/posts", api.Posts).Methods("GET")
	router.HandleFunc("/likes", api.Likes).Methods("GET")
	router.HandleFunc("/comments", api.Comments).Methods("GET")

	http.ListenAndServe(":8000", router)
}
