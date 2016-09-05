package writeas

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/writeas/impart"
	"io"
	"net/http"
	"time"
)

const (
	apiURL = "https://write.as/api"
)

type Client struct {
	baseURL string

	// Access token for the user making requests.
	token string
	// Client making requests to the API
	client *http.Client
}

// defaultHTTPTimeout is the default http.Client timeout.
const defaultHTTPTimeout = 10 * time.Second

func NewClient() *Client {
	return &Client{
		client:  &http.Client{Timeout: defaultHTTPTimeout},
		baseURL: apiURL,
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) get(path string, r interface{}) (*impart.Envelope, error) {
	method := "GET"
	if method != "GET" && method != "HEAD" {
		return nil, errors.New(fmt.Sprintf("Method %s not currently supported by library (only HEAD and GET).\n", method))
	}

	return c.request(method, path, nil, r)
}

func (c *Client) post(path string, data, r interface{}) (*impart.Envelope, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	return c.request("POST", path, b, r)
}

func (c *Client) request(method, path string, data io.Reader, result interface{}) (*impart.Envelope, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	r, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Create request: %v", err)
	}

	c.prepareRequest(r)
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Request: %v", err)
	}
	defer resp.Body.Close()

	env := &impart.Envelope{
		Code: resp.StatusCode,
	}
	if result != nil {
		env.Data = result
	}

	err = json.NewDecoder(resp.Body).Decode(&env)
	if err != nil {
		return nil, err
	}

	return env, nil
}

func (c *Client) prepareRequest(r *http.Request) {
	r.Header.Add("User-Agent", "go-writeas v1")
	r.Header.Add("Content-Type", "application/json")
	if c.token != "" {
		r.Header.Add("Authorization", "Token "+c.token)
	}
}
