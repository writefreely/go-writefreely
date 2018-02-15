package writeas

import (
	"testing"

	"fmt"
	"strings"
)

func TestCreatePost(t *testing.T) {
	wac := NewClient()
	p, err := wac.CreatePost(&PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		t.Errorf("Post create failed: %v", err)
		return
	}
	t.Logf("Post created: %+v", p)

	token := p.Token

	// Update post
	p, err = wac.UpdatePost(&PostParams{
		OwnedPostParams: OwnedPostParams{
			ID:    p.ID,
			Token: token,
		},
		Content: "Now it's been updated!",
	})
	if err != nil {
		t.Errorf("Post update failed: %v", err)
		return
	}
	t.Logf("Post updated: %+v", p)

	// Delete post
	err = wac.DeletePost(&PostParams{
		OwnedPostParams: OwnedPostParams{
			ID:    p.ID,
			Token: token,
		},
	})
	if err != nil {
		t.Errorf("Post delete failed: %v", err)
		return
	}
	t.Logf("Post deleted!")
}

func TestGetPost(t *testing.T) {
	wac := NewClient()

	res, err := wac.GetPost("zekk5r9apum6p")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		t.Logf("Post: %+v", res)
		if res.Content != "This is a post." {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}

	res, err = wac.GetPost("3psnxyhqxy3hq")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		if !strings.HasPrefix(res.Content, "                               Write.as Blog") {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}
}

func ExampleClient_CreatePost() {
	c := NewClient()

	// Publish a post
	p, err := c.CreatePost(&PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		fmt.Printf("Unable to create: %v", err)
		return
	}

	fmt.Printf("%s", p.Content)
	// Output: This is a post.
}
