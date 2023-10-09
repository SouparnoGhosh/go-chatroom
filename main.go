package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function that writes "Hello, World" to the response.
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Eripuka!")
	}

	// Register the handler function for the root ("/") route.
	http.HandleFunc("/", handler)

	// Start the web server on port 8080.
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}