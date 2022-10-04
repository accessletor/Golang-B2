package main

import (
	"fmt"
	"net/http"
)

// alamat port
var PORT = ":8080"

func main() {
	// handleFunc/route
	http.HandleFunc("/", greet)

	// menjalankan serve
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Web Server Golang"
	fmt.Fprint(w, msg)
}
