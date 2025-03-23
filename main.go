package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

const portNumber = ":8080"

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {

}
func renderTemplate(w http.ResponseseWriter, r *http.Request){
	parsedTemplate,_:=template.ParseFiles()
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
