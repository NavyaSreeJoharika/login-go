package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Login Page")
	http.HandleFunc("/hello", helloHandleFunc)

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontpage.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./sample.html")
	})

	http.ListenAndServe(":8080", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello")
}
