package writeas

import (
	"fmt"
	"net/http"
	"time"
)

type (
	// Post represents a published Write.as post, whether anonymous, owned by a
	// user, or part of a collection.
	Post struct {
		ID        string    `json:"id"`
		Slug      string    `json:"slug"`
		Token     string    `json:"token"`
		Font      string    `json:"appearance"`
		Language  *string   `json:"language"`
		RTL       *bool     `json:"rtl"`
		Listed    bool      `json:"listed"`
		Created   time.Time `json:"created"`
		Title     string    `json:"title"`
		Content   string    `json:"body"`
		Views     int64     `json:"views"`
		Tags      []string  `json:"tags"`
		Images    []string  `json:"images"`
		OwnerName string    `json:"owner,omitempty"`

		Collection *Collection `json:"collection,omitempty"`
	}

	// PostParams holds values for creating or updating a post.
	PostParams struct {
		// Parameters only for creating
		ID    string `json:"-"`
		Token string `json:"token,omitempty"`

		// Parameters for creating or updating
		Title    string  `json:"title,omitempty"`
		Content  string  `json:"body,omitempty"`
		Font     string  `json:"font,omitempty"`
		IsRTL    *bool   `json:"rtl,omitempty"`
		Language *string `json:"lang,omitempty"`

		Crosspost []map[string]string `json:"crosspost,omitempty"`
	}
)

func (c *Client) GetPost(id string) (*Post, error) {
	p := &Post{}
	env, err := c.get(fmt.Sprintf("/posts/%s", id), p)
	if err != nil {
		return nil, err
	}

	var ok bool
	if p, ok = env.Data.(*Post); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}
	status := env.Code

	if status == http.StatusOK {
		return p, nil
	} else if status == http.StatusNotFound {
		return nil, fmt.Errorf("Post not found.")
	} else if status == http.StatusGone {
		return nil, fmt.Errorf("Post unpublished.")
	} else {
		return nil, fmt.Errorf("Problem getting post: %s. %v\n", status, err)
	}
	return p, nil
}

func (c *Client) CreatePost(sp *PostParams) (*Post, error) {
	p := &Post{}
	env, err := c.post("/posts", sp, p)
	if err != nil {
		return nil, err
	}

	var ok bool
	if p, ok = env.Data.(*Post); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}

	status := env.Code
	if status == http.StatusCreated {
		return p, nil
	} else if status == http.StatusBadRequest {
		return nil, fmt.Errorf("Bad request: %s", env.ErrorMessage)
	} else {
		return nil, fmt.Errorf("Problem getting post: %s. %v\n", status, err)
	}
	return p, nil
}

func (c *Client) UpdatePost(sp *PostParams) (*Post, error) {
	p := &Post{}
	env, err := c.put(fmt.Sprintf("/posts/%s", sp.ID), sp, p)
	if err != nil {
		return nil, err
	}

	var ok bool
	if p, ok = env.Data.(*Post); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}

	status := env.Code
	if status == http.StatusOK {
		return p, nil
	} else if c.isNotLoggedIn(status) {
		return nil, fmt.Errorf("Not authenticated.")
	} else if status == http.StatusBadRequest {
		return nil, fmt.Errorf("Bad request: %s", env.ErrorMessage)
	} else {
		return nil, fmt.Errorf("Problem getting post: %s. %v\n", status, err)
	}
	return p, nil
}

func (c *Client) DeletePost(sp *PostParams) error {
	env, err := c.delete(fmt.Sprintf("/posts/%s", sp.ID), map[string]string{
		"token": sp.Token,
	})
	if err != nil {
		return err
	}

	status := env.Code
	if status == http.StatusNoContent {
		return nil
	} else if c.isNotLoggedIn(status) {
		return fmt.Errorf("Not authenticated.")
	} else if status == http.StatusBadRequest {
		return fmt.Errorf("Bad request: %s", env.ErrorMessage)
	}
	return fmt.Errorf("Problem getting post: %s. %v\n", status, err)
}
