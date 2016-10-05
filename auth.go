package writeas

import (
	"fmt"
	"net/http"
)

// LogIn authenticates a user with Write.as.
// See https://writeas.github.io/docs/#authenticate-a-user
func (c *Client) LogIn(username, pass string) (*AuthUser, error) {
	u := &AuthUser{}
	up := struct {
		Alias string `json:"alias"`
		Pass  string `json:"pass"`
	}{
		Alias: username,
		Pass:  pass,
	}

	env, err := c.post("/auth/login", up, u)
	if err != nil {
		return nil, err
	}

	var ok bool
	if u, ok = env.Data.(*AuthUser); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}

	status := env.Code
	if status == http.StatusOK {
		return u, nil
	} else if status == http.StatusBadRequest {
		return nil, fmt.Errorf("Bad request: %s", env.ErrorMessage)
	} else if status == http.StatusUnauthorized {
		return nil, fmt.Errorf("Incorrect password.")
	} else if status == http.StatusNotFound {
		return nil, fmt.Errorf("User does not exist.")
	} else if status == http.StatusTooManyRequests {
		return nil, fmt.Errorf("Stop repeatedly trying to log in.")
	} else {
		return nil, fmt.Errorf("Problem authenticating: %s. %v\n", status, err)
	}
	return u, nil
}

func (c *Client) isNotLoggedIn(code int) bool {
	if c.token == "" {
		return false
	}
	return code == http.StatusUnauthorized
}
