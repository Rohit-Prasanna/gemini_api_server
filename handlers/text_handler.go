package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gemini-api-server/utils" // Update the path as per your project structure
)

// TextRequest represents the input for text processing.
type TextRequest struct {
	Text string `json:"text"` // Input text for generation
}

// TextResponse represents the output of the text generation.
type TextResponse struct {
	GeneratedContent string `json:"generatedContent"` // Generated text response
	ErrorMessage     string `json:"error,omitempty"`  // Optional error message
}

// TextHandler processes text inputs using the Gemini API.
func TextHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON request body into TextRequest
	var req TextRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(TextResponse{ErrorMessage: "Invalid request body"})
		return
	}

	// Check if input text is empty
	if req.Text == "" {
		log.Println("Empty text input received")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(TextResponse{ErrorMessage: "Text input cannot be empty"})
		return
	}

	// Initialize the GenAI client
	client, err := utils.NewGenAIClient("AIzaSyDW-3zcMNqDAxTLTOUvHJqLvkjCHlZr6yY")
	if err != nil {
		log.Printf("Error creating GenAI client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(TextResponse{ErrorMessage: "Internal server error"})
		return
	}

	// Generate text using the GenAI client
	response, err := client.GenerateText(req.Text)
	if err != nil {
		log.Printf("Error generating text: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(TextResponse{ErrorMessage: "Failed to generate text"})
		return
	}

	// Successfully respond with the generated content
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TextResponse{GeneratedContent: response})
}
