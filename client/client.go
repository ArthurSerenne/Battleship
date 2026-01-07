package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		HttpClient: &http.Client{Timeout: 2 * time.Second},
	}
}

func (c *Client) Fire(url string, x, y int) (string, error) {
	payload := map[string]int{"x": x, "y": y}
	jsonData, _ := json.Marshal(payload)

	resp, err := c.HttpClient.Post(url+"/fire", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res["result"], nil
}
