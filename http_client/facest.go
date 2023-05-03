package facest

import (
	"net/http"
	"time"
)

const (
	BaseURL = "http://localhost:2012/v1"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewHttpClient(appKey string) *Client {
	return &Client{
		BaseURL: "",
		apiKey:  appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
