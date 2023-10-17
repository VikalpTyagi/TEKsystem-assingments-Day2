package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Student struct {
	Id    uint64
	Name  string
	Age   int
	Grade float32
}

var stuDatabase sync.Map

func main() {
	stu1 := Student{
		Id:    1234,
		Name:  "Vikalp Tyagi",
		Age:   24,
		Grade: 9.3,
	}
	stu2 := Student{
		Id:    5678,
		Name:  "Vishal",
		Age:   27,
		Grade: 9.5,
	}

	stuDatabase.Store(stu1.Id, stu1)
	stuDatabase.Store(stu2.Id, stu2)

	http.HandleFunc("/student", MidLogger(GetStudents))
	http.ListenAndServe(":8080", nil)
}

var ErrDataNotPresent = errors.New("Data not avilable in database")

func stuExisit(studentId uint64) (any, error) {
	data, ok := stuDatabase.Load(studentId)
	if !ok {
		return Student{}, ErrDataNotPresent
	}
	return data, nil
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	stuIdString := r.URL.Query().Get("students_Id")

	stuId, err := strconv.ParseUint(stuIdString, 10, 64)
	if err != nil {
		log.Println("Error while converting to Uint: ", err)
		errMap := map[string]string{"msg": "Input must be numerical only"}
		errJson, err := json.Marshal(errMap)
		if err != nil {
			log.Println("Error while parsing erro to json: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJson)
		return
	}
	studentData, err := stuExisit(stuId)
	if err != nil {
		log.Println("Error Id not present in database :", err)
		errMap := map[string]string{"msg": "Input must be numerical only"}
		errJson, err := json.Marshal(errMap)
		if err != nil {
			log.Println("Error while parsing data not exisit erro to json: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errJson)
		return
	}
	stuJson, err := json.Marshal(studentData)
	if err != nil {
		log.Println("Error while converting student data to json", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
		return
	}
	w.Write(stuJson)
}

func MidLogger(http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("I don't know how to writye this logic")
	}
}
