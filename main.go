package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsonProject/controller"
)

func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/app", controller.Index).Methods("GET")
	//router.HandleFunc("/app/get", controller.Index).Methods("GET")
	router.HandleFunc("/get/{degree}/{year}/{board}/{roll_id}", controller.FindResult).Methods("POST")

	http.ListenAndServe(":7010", router)
}
