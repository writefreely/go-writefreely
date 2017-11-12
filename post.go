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

	// OwnedPostParams are, together, fields only the original post author knows.
	OwnedPostParams struct {
		ID    string `json:"-"`
		Token string `json:"token,omitempty"`
	}

	// PostParams holds values for creating or updating a post.
	PostParams struct {
		// Parameters only for updating
		OwnedPostParams

		// Parameters for creating or updating
		Title    string  `json:"title,omitempty"`
		Content  string  `json:"body,omitempty"`
		Font     string  `json:"font,omitempty"`
		IsRTL    *bool   `json:"rtl,omitempty"`
		Language *string `json:"lang,omitempty"`

		Crosspost []map[string]string `json:"crosspost,omitempty"`

		// Parameters for collection posts
		Collection string `json:"-"`
	}

	// ClaimPostResult contains the post-specific result for a request to
	// associate a post to an account.
	ClaimPostResult struct {
		ID           string `json:"id,omitempty"`
		Code         int    `json:"code,omitempty"`
		ErrorMessage string `json:"error_msg,omitempty"`
		Post         *Post  `json:"post,omitempty"`
	}
)

// GetPost retrieves a published post, returning the Post and any error (in
// user-friendly form) that occurs. See
// https://developer.write.as/docs/api/#retrieve-a-post.
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
	}
	return nil, fmt.Errorf("Problem getting post: %s. %v\n", status, err)
}

// CreatePost publishes a new post, returning a user-friendly error if one comes
// up. See https://developer.write.as/docs/api/#publish-a-post.
func (c *Client) CreatePost(sp *PostParams) (*Post, error) {
	p := &Post{}
	endPre := ""
	if sp.Collection != "" {
		endPre = "/collections/" + sp.Collection
	}
	env, err := c.post(endPre+"/posts", sp, p)
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

// UpdatePost updates a published post with the given PostParams. See
// https://developer.write.as/docs/api/#update-a-post.
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
	}
	return nil, fmt.Errorf("Problem getting post: %s. %v\n", status, err)
}

// DeletePost permanently deletes a published post. See
// https://developer.write.as/docs/api/#delete-a-post.
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

// ClaimPosts associates anonymous posts with a user / account.
// https://developer.write.as/docs/api/#claim-posts.
func (c *Client) ClaimPosts(sp *[]OwnedPostParams) (*[]ClaimPostResult, error) {
	p := &[]ClaimPostResult{}
	env, err := c.put("/posts/claim", sp, p)
	if err != nil {
		return nil, err
	}

	var ok bool
	if p, ok = env.Data.(*[]ClaimPostResult); !ok {
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
	// TODO: does this also happen with moving posts?
	return p, nil
}
