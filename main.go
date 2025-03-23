package main

import (
	"fmt"
	"net/http"
)
func  Home(w http.ResponseWriter,r * http.Request){
	fmt.Fprintf(w,"This is the Home Page")

}
func About(w http.ResponseWriter,r *http.Request){
	sum,_:=AddValues(2,2)
	fmt.Fprintf(w,fmt.Sprintf("This is at the about page and 2+2 %d",sum))
}
func AddValues(x, y int)(int,error){

	
   return x+y,nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of Bytes Written: %d", n))

	})
	_ = http.ListenAndServe(":8080", nil)
}
