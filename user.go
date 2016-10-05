package writeas

import (
	"time"
)

type (
	// AuthUser represents a just-authenticated user. It contains information
	// that'll only be returned once (now) per user session.
	AuthUser struct {
		AccessToken string `json:"access_token,omitempty"`
		Password    string `json:"password,omitempty"`
		User        *User  `json:"user"`
	}

	// User represents a registered Write.as user.
	User struct {
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Created  time.Time `json:"created"`
	}
)
