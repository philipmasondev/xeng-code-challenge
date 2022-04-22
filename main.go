package main

import (
	"fmt"
	"net/http"
)

// port the server responds on
const portNumber = ":8080"

// HomePage serves the homepage to the request
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {

	http.HandleFunc("/", HomePage)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}
