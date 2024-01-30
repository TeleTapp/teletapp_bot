package api

import (
	"app/config"
	"app/logger"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"
)

type Client struct {
	baseURL string
	token   string
	logger  *zerolog.Logger
}

func NewClient() *Client {
	return &Client{
		baseURL: config.App.ApiBaseURL,
		token:   config.App.ApiKey,
		logger:  logger.NewLogger("ApiClient"),
	}
}

func (c *Client) post(ctx context.Context, path string, data interface{}) ([]byte, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+path, responseBody)
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
