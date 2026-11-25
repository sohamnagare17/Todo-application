package transport 

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"

	// "github.com/go-kit/kit/endpoint"
	// httptransport "github.com/go-kit/kit/transport/http"
)

type CreateTaskRequest struct {
    Title	string `json:"title"`
	Description string `json:"description"`
    Status string `json:"status"`
}

type CreateTaskResponse struct {
    Task interface{} `json:"task"`
    Err  string      `json:"error,omitempty"`
}

func DecodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("failed to decode request body: %w", err)
	}

	return req, nil
}

func EncodeCreateTaskResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(CreateTaskResponse)
	w.Header().Set("Content-Type", "application/json")	

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}