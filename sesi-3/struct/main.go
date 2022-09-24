package main

import (
	"fmt"
	"strings"
)

// Struct 1
type Employee struct {
	name     string
	age      int
	division string
}

// Struct 2
type Person struct {
	name string
	age  int
}

// Struct 3
type dataE struct {
	division string
	person   Person
}
func main() {
	// Struct (Giving value to struct)
	var employee Employee
	employee.name = "Asep San"
	employee.age = 22
	employee.division = "Software Engginer"
	fmt.Println("=============================================")
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println("=============================================")
	// Struct (Initializing struct)
	var employee1 = Employee{}
	employee1.name = "Asep"
	employee1.age = 22
	employee1.division = "Software Engginer"
	var employee2 = Employee{name: "Adi", age: 22, division: "Frontend Developer"}
	fmt.Printf("Employee1 %+v\n", employee1)
	fmt.Printf("Employee2 %+v\n", employee2)
	fmt.Println("=============================================")
	// Struct (Pointer to a struct)
	var dataEmployee1 = Employee{name: "Asep", age: 22, division: "Frontend Developer"}
	var dataEmployee2 *Employee = &dataEmployee1
	fmt.Println("Employee1 name:", dataEmployee1.name)
	fmt.Println("Employee2 name:", dataEmployee2.name)
	fmt.Println(strings.Repeat("#", 50))
	dataEmployee2.name = "Adi"
	fmt.Println("Employee1 name:", dataEmployee1.name)
	fmt.Println("Employee2 name:", dataEmployee2.name)
	fmt.Println("=============================================")
	// Struct (Embedded struct)
	var employ = dataE{}
	employ.person.name = "Asep"
	employ.person.age = 22
	employ.division = "Frontend Developer"

	fmt.Printf("%+v\n", employ)
	fmt.Println("=============================================")
	// Struct (Anonymous struct)
	// filed
	var employ1 = struct {
		person   Person
		division string
	}{}
	employ1.person = Person{name: "Asep", age: 22}
	employ1.division = "Frontend Developer"
	// No field
	var employ2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Adi", age: 22},
		division: "Fullstack",
	}

	fmt.Printf("Employee1: %+v\n", employ1)
	fmt.Printf("Employee2: %+v\n", employ2)
	fmt.Println("=============================================")

	// Struct (Slice of struct)
	var people = []Person{
		{name: "Asep", age: 22},
		{name: "Adi", age: 23},
		{name: "Agung", age: 24},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=============================================")
	// Struct (Slice of anonymous struct)
	var emp = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Asep", age: 22}, division: "Frontend Developer"},
		{person: Person{name: "Adi", age: 23}, division: "Backend Developer"},
		{person: Person{name: "Agung", age: 24}, division: "Fullstack Developer"},
	}

	for _, v := range emp {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=============================================")
}
