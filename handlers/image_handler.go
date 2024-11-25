package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gemini-api-server/utils"
)

// ImageRequest represents the input for image processing.
type ImageRequest struct {
	ImageData string `json:"imageData"` // Base64-encoded image
}

// ImageResponse represents the output of the image generation.
type ImageResponse struct {
	GeneratedContent string `json:"generatedContent"`
	ErrorMessage     string `json:"error,omitempty"` // Optional error message
}

// ImageHandler processes image inputs using the Gemini API.
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON request body into ImageRequest
	var req ImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ImageResponse{ErrorMessage: "Invalid request body"})
		return
	}

	// Initialize the GenAI client
	client, err := utils.NewGenAIClient("AIzaSyDW-3zcMNqDAxTLTOUvHJqLvkjCHlZr6yY")
	if err != nil {
		log.Printf("Error creating GenAI client: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ImageResponse{ErrorMessage: "Internal server error"})
		return
	}

	// Generate image content using the GenAI client
	response, err := client.GenerateImage(req.ImageData)
	if err != nil {
		log.Printf("Error generating image: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ImageResponse{ErrorMessage: "Failed to generate image content"})
		return
	}

	// Successfully respond with the generated content
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ImageResponse{GeneratedContent: response})
}
