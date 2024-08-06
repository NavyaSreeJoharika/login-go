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

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/sample", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./sample.html")
	})

	http.ListenAndServe(":8080", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello")
}
