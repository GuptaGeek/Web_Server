package main

import (
	"fmt"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")

}

const portNumber = ":8080"

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is at the about page and 2+2 %d", sum))
}

// addValues adds two integers and return the sum
func addValues(x, y int) (int, error) {
	return x + y, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
