package main

import (
	"crypto/md5" // Using a weak cryptographic algorithm
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func insecureHashing(input string) string {
	// Use a weak hashing function (MD5)
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Hard-coded credentials (Security Hotspot)
	secretKey := "hardCodedSasdasdecretKey123"
	username := "admin"

	// Log sensitive data (Security Hotspot)
	log.Printf("Username: %s, SecretKey: %s", username, secretKey)

	// Potential for SQL Injection if used in a query
	userInput := r.URL.Query().Get("input")

	// Complex logic
	if userInput == "" {
		userInput = "defaultUser"
		fmt.Fprintf(w, "Hello %s!", userInput)
	} else if userInput == "admin" {
		fmt.Fprintf(w, "Welcome admin!")
	} else {
		fmt.Fprintf(w, "Hello %s!", userInput)
	}
}

func main() {

	// Introduce duplicate code (Code Duplication)
	for i := 0; i < 3; i++ {
		fmt.Println("Starting server attempt:", i)
		startServer()
	}

	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 5000...")

	// Ignore error handling (Reliability Issue)
	_ = http.ListenAndServe(":5000", nil)

	// Introduce a bug: Infinite loop causing high CPU usage (New Bug)
	for {
		fmt.Println("This loop will run forever and cause high CPU usage")
	}
}

func startServer() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 5000...")
	_ = http.ListenAndServe(":5000", nil)
}
