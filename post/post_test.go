package post

import (
	"testing"

	"strings"
)

func TestGet(t *testing.T) {
	res := Get("3psnxyhqxy3hq")
	if !strings.HasPrefix(res, "                               Write.as Blog") {
		t.Errorf("Unexpected fetch results: %s\n", res)
	}
}
