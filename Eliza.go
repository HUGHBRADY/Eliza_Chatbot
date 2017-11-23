package main

import (
	"fmt"
	"net/http"
)

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("user-input")
	fmt.Fprintf(w, userInput)
}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/chat",userinputhandler)

	//http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}