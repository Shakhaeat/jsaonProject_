package model

import (
	"github.com/gchaincl/dotsql"
	"github.com/jsonProject/config"
	"github.com/jsonProject/entities"
)

type StudentModel struct{}

func (*StudentModel) FindAll() ([]entities.Student, error) {
	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		panic(err)
		//return , err
	}
	rows, err2 := db.Query("SELECT * from students")
	if err2 != nil {
		panic(err2)
		//return , err
	} else {
		var students []entities.Student
		for rows.Next() {
			var student entities.Student
			//rows.Scan(&student.ID, &student.RollID, &student.Name, &student.Cgpa, &student.Board)
			rows.Scan(&student.ID, &student.RollID, &student.Name, &student.FatherName, &student.MotherName,
				&student.Degree, &student.Year, &student.Board, &student.Institute, &student.Bangla,
				&student.GeneralMath, &student.English, &student.History, &student.CreatedAt)
			students = append(students, student)
			//students = append(students, student)
		}
		defer rows.Close()
		return students, nil
	}

}

func (*StudentModel) FindResult(roll_id int, degree string, board string, year int) (entities.Student, error) {
	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		panic(err)
		//return , err
	}
	// Loads queries from file
	dot, err2 := dotsql.LoadFromFile("queries.sql")
	if err2 != nil {
		panic(err)
	} else {

		// Run queries for create table
		//fmt.Println(student.RollID)
		//yearStr := strconv.FormatInt(student.Year, 10)
		//tab := student.Degree + yearStr
		//fmt.Println(dot)
		//rows, err4 := db.Query("Select * from "+tab+" WHERE roll_id=? AND board=?", student.RollID, student.Board)
		rows, err4 := dot.Query(db, "find-users-by-degree-year-board-roll", roll_id, degree, board, year)
		if err4 != nil {
			panic(err4)
			//log.Fatal(err)

		} else {
			//fmt.Println("Ayse")
			//var students []entities.Student
			var student entities.Student
			for rows.Next() {
				rows.Scan(&student.ID, &student.RollID, &student.Name, &student.FatherName, &student.MotherName,
					&student.Degree, &student.Year, &student.Board, &student.Institute, &student.Bangla,
					&student.GeneralMath, &student.English, &student.History, &student.CreatedAt)
				//students = append(students, student)
			}
			//	fmt.Println(students)
			return student, nil
		}
	}

}
