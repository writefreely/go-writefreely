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

func (c *Client) put(path string, data, r interface{}) (*impart.Envelope, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	return c.request("PUT", path, b, r)
}

func (c *Client) delete(path string, data map[string]string) (*impart.Envelope, error) {
	r, err := c.buildRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	for k, v := range data {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()

	return c.doRequest(r, nil)
}

func (c *Client) request(method, path string, data io.Reader, result interface{}) (*impart.Envelope, error) {
	r, err := c.buildRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	return c.doRequest(r, result)
}

func (c *Client) buildRequest(method, path string, data io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	r, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Create request: %v", err)
	}
	c.prepareRequest(r)

	return r, nil
}

func (c *Client) doRequest(r *http.Request, result interface{}) (*impart.Envelope, error) {
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

		err = json.NewDecoder(resp.Body).Decode(&env)
		if err != nil {
			return nil, err
		}
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
