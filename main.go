package main

import (
	"gemini-api-server/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5500", "https://rohit-prasanna.github.io/todo-list-withgo_backend/"}, // Replace with your frontend's URLs
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	// Apply CORS middleware to the router
	corsHandler := cors.New(corsOptions).Handler(r)

	// Define routes
	r.HandleFunc("/api/text", handlers.TextHandler).Methods("POST")
	r.HandleFunc("/api/image", handlers.ImageHandler).Methods("POST")

	// Start the server with CORS enabled
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
