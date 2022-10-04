package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// struct employee
type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

// inisialisasi variabel employess
var employees = []Employee{
	{ID: 1, Name: "Asep", Age: 22, Division: "IT"},
	{ID: 2, Name: "Saefuddin", Age: 22, Division: "Developer"},
	{ID: 3, Name: "Kun", Age: 22, Division: "Enginer"},
}

// penamaan port
var PORT = ":8080"

func main() {
	// membuat route /employees
	http.HandleFunc("/getEmployees", getEmployees)
	http.HandleFunc("/createEmployees", postEmployees)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// function get untuk mendapatkan value dari employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

// function post/create data employees
func postEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
