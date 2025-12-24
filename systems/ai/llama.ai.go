// Package ai is collection of AI api library
package ai

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"peregerine/systems/types/responses"
	"peregerine/configs"
)

// Function Generate is library to connect AI LLAMA client
func Generate(prompt string) (string, error) {
	request := responses.AiRequest{
		Model: "llama3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(request)

	client := &http.Client{
		Timeout: 120 * time.Second,
	}

	resp, err := client.Post(
		configs.LLAMAHost,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result responses.AiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}