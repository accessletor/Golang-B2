package main

import (
	"fmt"
	"strings"
)

func main() {

	//slice
	var buah = []string{"apple", "banana", "mango"}

	_ = buah

	// slice (make function)
	var sbuah2 = make([]string, 3)
	_ = sbuah2
	fmt.Printf("%#v\n", sbuah2)
	fmt.Println("======================================")
	//Slice (append function)
	var buah3 = make([]string, 3)
	buah3 = append(buah3, "apple", "banana", "mango")
	fmt.Printf("%#v\n", buah3)
	fmt.Println("======================================")

	//Slice (append function with ellipsis)
	var databuah1 = []string{"apple", "banana", "mango"}

	var databuah2 = []string{"durian", "pineapple", "starfruit"}

	databuah1 = append(databuah1, databuah2...)

	fmt.Printf("%#v\n", databuah1)
	fmt.Println("======================================")

	//Slice (copy function)
	var buah1 = []string{"apple", "banana", "mango"}

	var buah2 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(buah1, buah2)

	fmt.Println("Buah1 => ", buah1)
	fmt.Println("Buah2 => ", buah2)
	fmt.Println("Copied elements => ", nn)
	fmt.Println("======================================")

	// Slice (slicing)
	var dataBuah1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var dataBuah2 = dataBuah1[0:4]
	fmt.Printf("%#v\n", dataBuah2)

	var dataBuah3 = dataBuah1[0:]
	fmt.Printf("%#v\n", dataBuah3)

	var dataBuah4 = dataBuah1[:3]
	fmt.Printf("%#v\n", dataBuah4)

	var dataBuah5 = dataBuah1[:]
	fmt.Printf("%#v\n", dataBuah5)
	fmt.Println("======================================")

	//Slice (Combining slicing and append)
	var buahData = []string{"apple", "banana", "mango", "durian", "pineapple"}
	buahData = append(buahData[:3], "rambutan")
	fmt.Printf("%#v\n", buahData)

	fmt.Println("======================================")

	//Slice (Backing array)
	var buahData1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var buahData2 = buahData1[2:4]

	buahData2[0] = "rambutan"

	fmt.Println("buah1 => ", buahData1)
	fmt.Println("buah2 => ", buahData2)
	fmt.Println("======================================")

	//Slice (Cap function)
	var dBuah1 = []string{"apple", "banana", "mango", "durian"}

	fmt.Println("buah1 cap: ", cap(dBuah1))
	fmt.Println("buah1 len: ", len(dBuah1))

	fmt.Println(strings.Repeat("#", 20))

	var dBuah2 = dBuah1[0:3]

	fmt.Println("buah1 cap: ", cap(dBuah2))
	fmt.Println("buah1 len: ", len(dBuah2))

	fmt.Println(strings.Repeat("#", 20))

	var dBuah3 = dBuah1[1:]

	fmt.Println("buah1 cap: ", cap(dBuah3))
	fmt.Println("buah1 len: ", len(dBuah3))
	fmt.Println("======================================")

	//Slice (Creating a new backing array)
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars: ", cars)
	fmt.Println("newcars: ", newCars)
	fmt.Println("======================================")
}
