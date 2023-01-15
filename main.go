package main

import (
	"fmt"
	"net/http"
)

const portNUmber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server starting on port %s", portNUmber))
	_ = http.ListenAndServe(portNUmber, nil)

}
