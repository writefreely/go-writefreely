package writeas

import (
	"testing"

	"strings"
)

func TestGet(t *testing.T) {
	wac := NewClient("")

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
