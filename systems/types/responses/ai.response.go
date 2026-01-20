// Package responses is bundle of function to get base response
package responses

// Type AiRequest is for request to access AI API
type AiRequest struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool `json:"stream"`
}

// Type AiResponse is for base response from AI API
type AiResponse struct {
	Response string `json:"response"`
	Done bool `json:"done"`
}
