package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		// Respond with a simple message
		fmt.Fprintf(w, "Hello, %s! Welcome to your first Go HTTP server.", name)
	})

	// Define the handler function for About
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the About page of your HTTP server.")
	})

	// Start the HTTP server on port 8080
	fmt.Println("Starting server on http://localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Log an error if the server fails to start
		fmt.Println("Error starting server:", err)
	}
}
