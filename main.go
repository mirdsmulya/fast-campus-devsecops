package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "Mirdan"
	}
	fmt.Fprintf(w, "Hello World from %s!", name)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 5000...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
