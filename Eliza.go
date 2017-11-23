// Eliza Chatbot by Hugh Brady 2017

package main

import (
	"math/rand"
	"time"
	"fmt"
	"regexp"
	"strings"
)

func elizaResponse(input string) string {

	// Array of strings 
	responses := []string{
		"Why do you say that?",
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	
	return responses[rand.Intn(len(responses))]
}

func main() {
	dir := http.Dir("webpage")
	fs := http.FileServer(http.Dir(dir)

	http.Handle("/", fs)
	http.HandleFunc("/user-input", inputHandler)
	http.ListenAndServe(":8080", nil)
}