package writeas

import (
	"fmt"
	"testing"
)

func TestGetCollection(t *testing.T) {
	wac := NewClient()

	res, err := wac.GetCollection("blog")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		t.Logf("Post: %+v", res)
		if res.Title != "write.as" {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}
}

func ExampleClient_GetCollection() {
	c := NewClient()
	coll, err := c.GetCollection("blog")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%+v", coll)
}
