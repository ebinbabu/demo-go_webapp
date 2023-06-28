package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the directory where the index.html file is located
	dir := "./"

	// Create a file server handler for serving static files
	fileServer := http.FileServer(http.Dir(dir))

	// Register the file server handler for the root URL path "/"
	http.Handle("/", fileServer)

	// Start the HTTP server on port 8080
	port := ":80"
	fmt.Printf("Server listening on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
