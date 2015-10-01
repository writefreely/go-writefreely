package writeas

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiURL = "https://write.as/api"
)

type API struct {
	BaseURL string
}

// defaultHTTPTimeout is the default http.Client timeout.
const defaultHTTPTimeout = 10 * time.Second

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

func GetAPI() *API {
	return &API{apiURL}
}

func (a API) Call(method, path string) (int, string, error) {
	if method != "GET" && method != "HEAD" {
		return 0, "", errors.New(fmt.Sprintf("Method %s not currently supported by library (only HEAD and GET).\n", method))
	}

	r, _ := http.NewRequest(method, fmt.Sprintf("%s%s", a.BaseURL, path), nil)
	r.Header.Add("User-Agent", "writeas-go v1")

	resp, err := httpClient.Do(r)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}

	return resp.StatusCode, string(content), nil
}
