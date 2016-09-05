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
		ID          string    `json:"id"`
		Slug        string    `json:"slug"`
		ModifyToken string    `json:"token"`
		Font        string    `json:"appearance"`
		Language    *string   `json:"language"`
		RTL         *bool     `json:"rtl"`
		Listed      bool      `json:"listed"`
		Created     time.Time `json:"created"`
		Title       string    `json:"title"`
		Content     string    `json:"body"`
		Views       int64     `json:"views"`
		Tags        []string  `json:"tags"`
		Images      []string  `json:"images"`
		OwnerName   string    `json:"owner,omitempty"`

		Collection *Collection `json:"collection,omitempty"`
	}

	// PostParams holds values for creating or updating a post.
	PostParams struct {
		Title    string  `json:"title"`
		Content  string  `json:"body"`
		Font     string  `json:"font"`
		IsRTL    *bool   `json:"rtl"`
		Language *string `json:"lang"`

		Crosspost []map[string]string `json:"crosspost"`
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
