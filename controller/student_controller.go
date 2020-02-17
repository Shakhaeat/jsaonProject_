package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jsonProject/model"
)

// func Index(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var studentModel model.StudentModel
// 	students, _ := studentModel.FindAll()
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(students)

// }

func FindResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	params := mux.Vars(r)
	rollID, _ := strconv.Atoi(params["roll_id"])
	degree := params["degree"]
	board := params["board"]
	year, _ := strconv.Atoi(params["year"])
	//fmt.Println(rollID)
	var studentModel model.StudentModel
	student, err := studentModel.FindResult(rollID, degree, board, year)
	if err != nil {
		panic(err)
	}
	//w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)

}
