package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://gettestmail.com/api"

type GetTestMailClient struct {
	APIKey string
	http   *http.Client
}

func NewGetTestMailClient(apiKey string) GetTestMailClient {
	return GetTestMailClient{
		APIKey: apiKey,
		http:   &http.Client{
			// Timeout: 15 * time.Minute, // needs to be set the same as the expiration time of the email address
		},
	}
}

func (c GetTestMailClient) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("X-API-Key", c.APIKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode > 399 {
		var problem Problem
		err = json.Unmarshal(body, &problem)
		if err != nil {
			return nil, errors.New("failed to parse response")
		}
		return nil, fmt.Errorf("an error occurred: %v", problem.Detail)
	}

	return body, nil
}

func (c GetTestMailClient) CreateNew(ctx context.Context) (*GetTestMail, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/gettestmail", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to createNew getTestMail: %w", err)
	}

	var getTestMail GetTestMail
	err = json.Unmarshal(body, &getTestMail)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &getTestMail, nil
}

func (c GetTestMailClient) WaitForMessage(ctx context.Context, id string) (*GetTestMail, error) {
	url := fmt.Sprintf("%s/gettestmail/%s", baseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("empty response body")
	}

	var getTestMail GetTestMail
	err = json.Unmarshal(body, &getTestMail)
	if err != nil {
		return nil, err
	}

	if getTestMail.Message != nil {
		return &getTestMail, nil
	}

	return nil, nil
}
