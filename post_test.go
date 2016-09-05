package writeas

import (
	"testing"

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
	} else {
		t.Logf("Post created: %+v", p)
	}
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
