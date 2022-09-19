package main

import "fmt"

func main() {
	//belajar hello
	fmt.Println("Hello")
	fmt.Println("world")

	fmt.Println("Hello World")
	fmt.Println("Hello", "World")

	fmt.Print("Hello World\n")
	fmt.Print("Hello", " World\n")
	fmt.Print("Hello", " ", "World\n")

	fmt.Println("==========================")

	//belajar variable dengan data type
	var nama string
	var umur int = 19

	nama = "Asep Saefuddin"
	umur = 22

	fmt.Println("Nama Saya", nama)
	fmt.Println("umur Saya", umur)
	fmt.Println("==========================")

	//belajar variable tanpa data type
	var nama2 = "Jane"
	var umur2 = 25

	fmt.Printf("%T,%T\n", nama2, umur2)
	fmt.Println("==========================")

	//belajar short declaration
	namaSaya := "Ariel"
	umurSaya := 15

	fmt.Printf("%T,%T\n", namaSaya, umurSaya)
	fmt.Println("==========================")

	//belajar multiple variable declarations
	var coba1, coba2, coba3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 3

	fmt.Println(coba1, coba2, coba3)
	fmt.Println(first, second, third)

	var name1, age1, address1 = "Aurora", 16, "Castle"
	data1, data2, data3 := 1, "Apple", 3.5

	fmt.Println(name1, age1, address1)
	fmt.Println(data1, data2, data3)
	fmt.Println("==========================")

	//belajar underscore variable
	//untuk menghindari error jika suatu variable tidak digunakan
	var firstVariable string
	var name2, age2, address2 = "Asep", 22, "Nazarick"

	_, _, _, _ = firstVariable, name2, age2, address2
	fmt.Printf("test underscore variabel > %T\n", firstVariable)
	fmt.Printf("Halo nama saya %s, umur saya adalah %d tahun, saya tinggal di %s\n", name2, age2, address2)
	fmt.Println("==========================")

	//Data type Number (Integer)
	var num1 uint8 = 89
	var num2 int8 = -12
	fmt.Printf("tipe data num1: %T\n", num1)
	fmt.Printf("tipe data num2: %T\n", num2)
	fmt.Println("==========================")

	//Data type Numer (Float)
	var decimal1 float32 = 3.63
	fmt.Printf("decimal number: %f\n", decimal1)
	fmt.Printf("decimal number: %.3f\n", decimal1)
	fmt.Println("==========================")

	//data type boolean
	var cond bool = true
	fmt.Printf("is it permitted? %t\n", cond)
	fmt.Println("==========================")

	//data type string
	var message = `Selamat datang di "Hacktiv8". Salam kenal :). Mari belajar "Scalable Web Service With Go"`
	fmt.Println(message)
	fmt.Println("==========================")

	//constant variable
	const full_name string = "John Doe"
	fmt.Printf("Hello %s\n", full_name)
	fmt.Println("==========================")

	//Operators
	var value = 2 + 2*3
	fmt.Println(value)

	var value2 = (2 + 2) * 3
	fmt.Println(value2)

	fmt.Println("==========================")

	//Relational Operators

	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var forthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("forth condition:", forthCondition)

	fmt.Println("==========================")

	//Logical Operators
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("Wrong && Right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("Wrong || Right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)

}