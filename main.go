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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no port is specified
	}

	// Start the server with the CORS handler on the correct port
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
