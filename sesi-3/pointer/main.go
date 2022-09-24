package main

import (
	"fmt"
	"strings"
)

func main() {
	//Pointer
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber
	//Pointer (memory address)
	var numberOne int = 4
	var numberTwo *int = &numberOne
	fmt.Println("===============================================")
	fmt.Println("Value number One : ", numberOne)
	fmt.Println("Address memory numberOne : ", &numberOne)

	fmt.Println("Value numberTwo : ", *numberTwo)
	fmt.Println("Address memory numberTwo : ", &numberTwo)
	fmt.Println("===============================================")
	// Pointer (Changing value through pointer)
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (memori address) : ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 50), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (memori address) : ", secondPerson)
	fmt.Println("===============================================")
	// Pointer (Pointer as a parameter)
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
	fmt.Println("===============================================")
}

// func changeValue
func changeValue(number *int) {
	*number = 20
}
