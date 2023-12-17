package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to version 2 of the great lenslocked</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting Server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
