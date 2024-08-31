// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	name := os.Getenv("NAME")
// 	if name == "" {
// 		name = "Mirdan"
// 	}
// 	fmt.Fprintf(w, "Hello World from %s!", name)
// }

// func main() {
// 	http.HandleFunc("/", helloHandler)
// 	fmt.Println("Server starting on port 5000...")
// 	err := http.ListenAndServe(":5000", nil)
// 	if err != nil {
// 		fmt.Println("Error starting server: ", err)
// 	}
// }

package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Hard-coded secret key (Security Hotspot)
	secretKey := "superSecretKey123"

	// Use query parameter without validation (Code Smell / Security Hotspot)
	name := r.URL.Query().Get("name")
	if name == "" {
		name = os.Getenv("NAME")
		if name == "" {
			name = "Mirdan"
		}
	}

	// Print the secret key to the console (Security Hotspot)
	fmt.Println("Using secret key: ", secretKey)

	// Potential XSS vulnerability by reflecting user input
	fmt.Fprintf(w, "Hello World from %s!", name)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 5000...")

	// Ignore error returned by ListenAndServe (Code Smell)
	_ = http.ListenAndServe(":5000", nil)
}
