package main

import (
	"fmt"
	"net/http"
)

const portNUmber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Number of bytes: %d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server starting on port %s", portNUmber))
	_ = http.ListenAndServe(portNUmber, nil)

}
