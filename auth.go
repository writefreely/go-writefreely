package writeas

import (
	"net/http"
)

func (c *Client) isNotLoggedIn(code int) bool {
	if c.token == "" {
		return false
	}
	return code == http.StatusUnauthorized
}
