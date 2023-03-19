package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/about", aboutPage)
	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World from Go! Main Page")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Welcome to the about page!</h1><p>This is a fancy web page.</p></body></html>")
}