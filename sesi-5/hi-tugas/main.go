package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)

func main() {
	fmt.Println("===============================")
	var username string
	var username2 string
	fmt.Println("Create Account")
	fmt.Print("Input Username : ")
	fmt.Scanln(&username)
	fmt.Print("Input Password : ")
	password, _ := gopass.GetPasswdMasked()
	fmt.Println("===============================")

	fmt.Println("Log In")
	fmt.Print("Input username : ")
	fmt.Scanln(&username2)
	fmt.Print("Input Password : ")
	password2, _ := gopass.GetPasswdMasked()
	if valid, err := validasi(username, username2, password, password2); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}

	fmt.Println("===============================")
}

func validasi(username string, username2 string, password []byte, password2 []byte) (string, error) {
	if username != username2 {
		return "", errors.New("Username/password invalid")
	}

	return "Username & password valid", nil
}
