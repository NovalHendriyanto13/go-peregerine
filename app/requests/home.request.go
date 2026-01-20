// Package requests is group of request for controller
package requests

// Type HomeCreateRequest is type for Home to create transaction
type HomeCreateRequest struct{
	Message string `json:"message"`
}