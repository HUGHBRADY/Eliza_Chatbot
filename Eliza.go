package main

import (
	"fmt"
	"net/http"
)

func elizaBrain(userInput string) string {
	return "how now brown cow"
}

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("user-input")
	// pass user input into Eliza's brain
	resp := elizaBrain(userInput)
	 
	fmt.Fprintf(w, resp)
}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/chat", userinputhandler)

	//http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}