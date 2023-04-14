package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(writer, request)
	})
}
