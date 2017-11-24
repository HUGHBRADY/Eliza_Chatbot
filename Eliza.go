// ========= Eliza Chatbot Project =========
// ========= Hugh Brady  G00338260 =========

package main

import (
	"fmt"
	"bufio"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"net/http"
)

// Code adapted from https://github.com/data-representation/eliza

// ========= STRUCTS =========

// A struct consisting of a regular expression and an array of strings
// Eliza's response will be picked randomly from the array
type Replacer struct {
	rExpression  *regexp.Regexp
	replacements []string
}

// This struct contains arrays of the replacer struct, which will be used to return Eliza's response 
type Eliza struct {
	responses     []Replacer
	substitutions []Replacer
}

// ========= FUNCTIONS =========

// Populates an array of Replacers using info from a text file.
func ReadReplacersFromFile(path string) []Replacer {
	
	file, err := os.Open(path)						// Open the file
	if err != nil {									// If there's an error opening the file,
		log.Fatal(err)								// log a fatal error
	}
	defer file.Close()								// Close file when done

	var replacers []Replacer						// Create an empty array of Replacer structs

	// Read the file line by line
	for scanner, readRegExpression := bufio.NewScanner(file), false; scanner.Scan(); {
		switch line := scanner.Text(); {
				
			case strings.HasPrefix(line, "#"):		// If the line starts with a # then it is a comment 
				
			case len(line) == 0:					// Indicates a blank line
				readRegExpression = false			// Set var to false before we come to a new section 

			case readRegExpression == false:		// the first line that isn't blank is the reg expression
				replacers = append(replacers, Replacer{rExpression: regexp.MustCompile(line)})
				readRegExpression = true			// Set to true
			
			default:
				// Otherwise read a replacement and add it to the last replacer.
				replacers[len(replacers)-1].replacements = append(replacers[len(replacers)-1].replacements, line)
		}
	}
	
	return replacers								// Return the replacers array
}

// Takes in user input and returns a response
func (me *Eliza) RespondTo(input string) string {
	// Look for a suitable response
	for _, response := range me.responses {
		// Check if user input matches a regular expression
		if matches := response.rExpression.FindStringSubmatch(input); matches != nil {		
			
			output := response.replacements[rand.Intn(len(response.replacements))]	// Select a random response from the array
			
			captured := regexp.MustCompile(`[\s,.?!]+`)				// Tokenise the captured groups using the following regular expression
		
			for m, match := range matches[1:] {						// Fill the response with each captured group from the input

				tokens := captured.Split(match, -1)					// Split captured into tokens
				
				for t, token := range tokens {						// Loop through the tokens
					
					for _, substitution := range me.substitutions {	// Loop through the substitutions
						if substitution.rExpression.MatchString(token) {
							// If the original matches, replace the token with one of the replacements
							tokens[t] = substitution.replacements[rand.Intn(len(substitution.replacements))]
							break
						}
					}
				}
				// Replace $1 with the first match, $2 with the second and so on
				output = strings.Replace(output, "$"+strconv.Itoa(m+1), strings.Join(tokens, " "), -1)
			}
			return output
		}
	}
	// generic response
	return "Stop talking gibberish."
}

// Gives Eliza access to the data files
func ElizaFromFiles(responsePath string, substitutionPath string) Eliza {
	eliza := Eliza{}												// Initialize eliza as the struct

	eliza.responses = ReadReplacersFromFile(responsePath)			// Populate the replacer array "responses" with data in responses.txt
	eliza.substitutions = ReadReplacersFromFile(substitutionPath)	// Same with substitutions

	return eliza													// Returns populated instance of the Eliza struct
}

// ========= Handlers =========

// Handles /chat requests. Is basically the mediator between Eliza and the html
func userinputhandler(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("user-input")								// Takes in user input from jQuery in js file 
	
	eliza := ElizaFromFiles("data/responses.txt", "data/substitutions.txt")		// Create a new instance of Eliza.

	fmt.Fprintf(w, eliza.RespondTo(userInput))									// Returns Eliza's response to the js
}

func main() {
	rand.Seed(time.Now().UnixNano()) 				// Seed the rand package with the current time.

	// File serving
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/chat", userinputhandler)		// Request handler
	http.ListenAndServe(":8080", nil)
}