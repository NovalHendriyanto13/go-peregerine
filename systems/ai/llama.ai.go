// Package ai is collection of AI api library
package ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"peregerine/systems/types/responses"
	"peregerine/configs"
)

var client = &http.Client{
    Timeout: 120 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 100,
        IdleConnTimeout:     90 * time.Second,
    },
}

// Generate is library to connect AI LLAMA client
func Generate(prompt string) (string, error) {
	request := responses.AiRequest{
		Model: "llama3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(request)

	// client := &http.Client{
	// 	Timeout: 120 * time.Second,
	// }

	resp, err := client.Post(
		configs.LLAMAHost + "/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() } ()

	var result responses.AiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}

func StreamGenerate(prompt string, out chan<- string) error {
	request := responses.AiRequest{
		Model: "llama3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(request)

	// client := &http.Client{
	// 	Timeout: 120 * time.Second,
	// }

	resp, err := client.Post(
		configs.LLAMAHost + "/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() } ()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		var chunk responses.AiResponse
		line := scanner.Bytes()

		if err := json.Unmarshal(line, &chunk); err != nil {
			continue
		}

		if chunk.Response != "" {
			out <- chunk.Response
		}

		if chunk.Done {
			break
		}
	}

	return nil
}