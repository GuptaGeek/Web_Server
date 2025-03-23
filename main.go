package main

import (
	"errors"
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
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0 ")
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))

}
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
