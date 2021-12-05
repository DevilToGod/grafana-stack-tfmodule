package gapi

import (
	"bytes"
	"encoding/json"
)

type CreateStackRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Region string `json:"region"`
}

type CreateStackResponse struct {
	// ID field only returned after Grafana v7.
	URL       string `json:"url"`
	ClusterID string `json:"clusterId"`
}

// CreateAPIKey creates a new Grafana API key.
func (c *Client) CreateStack(request CreateStackRequest) (CreateStackResponse, error) {
	response := CreateStackResponse{}

	data, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	response, err = c.request("POST", "/api/instances", nil, bytes.NewBuffer(data), &response)
	return response, err
}
