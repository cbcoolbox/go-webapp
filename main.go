package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("LOADED HOME")
	fmt.Fprintf(w, "this is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the about page")
}

func Divide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the divide page: ")
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}

	result := x / y

	return result, nil
}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	http.HandleFunc("/divide", Divide)

	http.ListenAndServe(":8080", nil)
}
