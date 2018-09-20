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
		t.Logf("Collection: %+v", res)
		if res.Title != "write.as" {
			t.Errorf("Unexpected fetch results: %+v\n", res)
		}
	}
}

func TestGetCollectionPosts(t *testing.T) {
	wac := NewClient()

	res, err := wac.GetCollectionPosts("blog")
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		if len(*res) == 0 {
			t.Errorf("No posts returned!")
		}
	}
}

func TestGetUserCollections(t *testing.T) {
	wac := NewDevClient()
	_, err := wac.LogIn("demo", "demo")
	if err != nil {
		t.Fatalf("Unable to log in: %v", err)
	}
	defer wac.LogOut()

	res, err := wac.GetUserCollections()
	if err != nil {
		t.Errorf("Unexpected fetch results: %+v, err: %v\n", res, err)
	} else {
		t.Logf("User collections: %+v", res)
		if len(*res) == 0 {
			t.Errorf("No collections returned!")
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
	fmt.Printf("%s", coll.Title)
	// Output: write.as
}
