package main

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade int
}

type InsertStudent struct {
	Name  string
	Age   int
	Grade int
}

var db *sqlx.DB

func main() {
	var err error
	connStr := "postgres://ts:ts@localhost:5432/postgres?sslmode=disable"
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Database Connected!")
	}

	// GET ALL HERE
	students, err := GetStudents()
	if err != nil {
		panic(err)
	}
	fmt.Println(students)

	// GET HERE
	student, err := GetStudent(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(student)

	// INSERT HERE
	/*studentInsert := InsertStudent{"Thanyawat", 24, 20}
	res, err := AddStudent(studentInsert)
	if err != nil {
		fmt.Println(res)
		panic(err)
	} else {
		fmt.Println(res)
	}*/

	//DELETE HERE
	/*deleteRes, err := DeleteStudent(4)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteRes)*/

	// UPDATE HERE
	/*studentUpdate := InsertStudent{"TS", 20, 210}
	res, err := UpdateStudent(5, studentUpdate)
	if err != nil {
		fmt.Println(res)
		panic(err)
	} else {
		fmt.Println(res)
	}*/
}

func GetStudents() ([]Student, error) {
	queryStr := "select id,name,age,grade from student"
	students := []Student{}
	err := db.Select(&students, queryStr)
	if err != nil {
		return nil, err
	}
	return students, nil
}

func GetStudent(id int) (*Student, error) {
	queryStr := "select id,name,age,grade from student WHERE id = $1"
	student := Student{}
	err := db.Get(&student, queryStr, id)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func AddStudent(student InsertStudent) (string, error) {
	queryStr := "insert into student (name, age, grade) values ($1,$2,$3)"
	tx, err := db.Begin()
	if err != nil {
		return "INSERT Failure", err
	}
	res, err := tx.Exec(queryStr, student.Name, student.Age, student.Grade)
	if err != nil {
		return "INSERT Failure", err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return "Row effected Failure", err
	}

	if affected <= 0 {
		return "Row effected is zero", errors.New("Cannot Insert")
	}
	err = tx.Commit()
	if err != nil {
		return "Commit Failure", err
	}

	return "Successfully!", nil
}

func DeleteStudent(id int) (string, error) {
	queryStr := "delete from student where id=$1"
	tx, err := db.Begin()
	if err != nil {
		return "DELETE Failure", err
	}
	res, err := tx.Exec(queryStr, id)
	if err != nil {
		return "DELETE Failure", err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return "Row effected Failure", err
	}

	if affected <= 0 {
		return "Row effected is zero", errors.New("Cannot Insert")
	}
	err = tx.Commit()
	if err != nil {
		return "Commit Failure", err
	}

	return "Successfully!", nil
}

func UpdateStudent(id int, student InsertStudent) (string, error) {
	queryStr := "UPDATE student SET (name, age, grade) = ($2, $3, $4) where id=$1"
	tx, err := db.Begin()
	if err != nil {
		return "DELETE Failure", err
	}
	res, err := tx.Exec(queryStr, id, student.Name, student.Age, student.Grade)
	if err != nil {
		return "DELETE Failure", err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return "Row effected Failure", err
	}

	if affected <= 0 {
		return "Row effected is zero", errors.New("Cannot Insert")
	}
	err = tx.Commit()
	if err != nil {
		return "Commit Failure", err
	}
	return "Successfully!", nil
}
