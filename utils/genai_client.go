package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GenAIClient wraps the Generative AI client.
type GenAIClient struct {
	client *genai.Client
}

// NewGenAIClient initializes and returns a new GenAI client.
func NewGenAIClient(apiKey string) (*GenAIClient, error) {
	ctx := context.Background()

	// Initialize the Generative AI client with the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Generative AI client: %v", err)
	}

	return &GenAIClient{client: client}, nil
}

// GenerateText generates text using the Generative AI client.
func (c *GenAIClient) GenerateText(prompt string) (string, error) {
	ctx := context.Background()

	// Select the generative model
	model := c.client.GenerativeModel("gemini-1.5-flash")

	// Generate content using the model
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("error generating text: %v", err)
	}

	// Initialize result
	var result string

	// Extract response parts and concatenate
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				// Use fmt.Sprintf to ensure any type is converted to a string
				result += fmt.Sprintf("%v\n", part)
			}
		}
	}

	// Check if result is empty
	if result == "" {
		return "", fmt.Errorf("no valid content returned in response")
	}

	return result, nil
}

func (c *GenAIClient) GenerateImage(base64Image string) (string, error) {
	ctx := context.Background()

	// Select the generative model
	model := c.client.GenerativeModel("gemini-1.5-flash")

	// Decode the Base64 image into raw byte data
	imgData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %v", err)
	}

	// Generate content using the model
	resp, err := model.GenerateContent(ctx, genai.ImageData("jpeg", imgData))
	if err != nil {
		return "", fmt.Errorf("error generating image content: %v", err)
	}

	// Initialize result
	var result string

	// Extract response parts and concatenate
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				// Use fmt.Sprintf to ensure any type is converted to a string
				result += fmt.Sprintf("%v\n", part)
			}
		}
	}

	// Check if result is empty
	if result == "" {
		return "", fmt.Errorf("no valid content returned in response")
	}

	return result, nil
}

// decodeBase64 decodes a Base64 string into raw byte data.
func decodeBase64(base64String string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return data, nil
}
