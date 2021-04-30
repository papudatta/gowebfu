package main

import (
	"fmt"
	"net/http"
	"errors"
)

func Home(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is the home page")
	}


func About(w http.ResponseWriter, r *http.Request) {
	sum := addvalues (2, 3) 
	fmt.Fprint(w, fmt.Sprintf("This is the about page and 2 + 3 = %d", sum))
	}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32
	//x, y = 121.00, 10.00
	x, y = 121.00, 0.00
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	return x/y, nil
}

func addvalues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe("0.0.0.0:9090", nil)
}
