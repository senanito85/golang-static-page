package main

import (
	"log"
	"net/http"
)

func main() {
	// Log that the server is starting
	log.Println("Starting server on port 8080...")

	// Redirect the root path ("/") to the "/home" path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusFound) // StatusFound = 302
	})

	// Serve the "home.html" file at the "/home" path
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html") // Ensure the file path is correct
	})

	// Serve the "info.html" file at the "/info" path
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./info.html") // Ensure the file path is correct
	})

	// Serve the "contacts.html" file at the "/contacts" path
	http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./contacts.html") // Ensure the file path is correct
	})

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
