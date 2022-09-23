package main

import (
	"fmt"
)

func main() {
	//Loopings (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka ", i)
	}
	fmt.Println("======================================")

	//Loopings (second way of looping)
	var second = 0
	for second < 3 {
		fmt.Println("Angka ", second)
		second++
	}
	fmt.Println("======================================")

	//Loopings (third way of looping)
	var thrid = 0
	for {
		fmt.Println("Angka ", thrid)

		thrid++
		if thrid == 3 {
			break
		}
	}
	fmt.Println("======================================")

	//Loopings (break and continue keywords)
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka ", i)
	}
	fmt.Println("======================================")

	//Loopings (Nested Looping)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
	fmt.Println("======================================")

	//Loopings (Label)
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := i; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
