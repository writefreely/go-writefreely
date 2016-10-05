package writeas

import (
	"fmt"
	"net/http"
)

// Collection represents a collection of posts. Blogs are a type of collection
// on Write.as.
type Collection struct {
	Alias       string `json:"alias"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StyleSheet  string `json:"style_sheet"`
	Private     bool   `json:"private"`
	Views       int64  `json:"views"`
	Domain      string `json:"domain,omitempty"`
	Email       string `json:"email,omitempty"`

	TotalPosts int `json:"total_posts"`
}

// GetCollection retrieves a collection, returning the Collection and any error
// (in user-friendly form) that occurs. See
// https://writeas.github.io/docs/#retrieve-a-collection
func (c *Client) GetCollection(alias string) (*Collection, error) {
	coll := &Collection{}
	env, err := c.get(fmt.Sprintf("/collections/%s", alias), coll)
	if err != nil {
		return nil, err
	}

	var ok bool
	if coll, ok = env.Data.(*Collection); !ok {
		return nil, fmt.Errorf("Wrong data returned from API.")
	}
	status := env.Code

	if status == http.StatusOK {
		return coll, nil
	} else if status == http.StatusNotFound {
		return nil, fmt.Errorf("Collection not found.")
	} else {
		return nil, fmt.Errorf("Problem getting collection: %s. %v\n", status, err)
	}
	return coll, nil
}
