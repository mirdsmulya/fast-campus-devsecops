package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get the NAME environment variable
	name := os.Getenv("NAME")
	if name == "" {
		name = "Mirdan" // Default value if NAME is not set
	}

	// Respond with a message
	fmt.Fprintf(w, "Hello World from %s!", name)
}

func main() {
	// Register the handler function for the root path
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server on port 5000
	fmt.Println("Server starting on port 5000...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
