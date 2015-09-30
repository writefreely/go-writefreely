package post

import (
	"fmt"
	writeas "github.com/writeas/writeas-go"
	"net/http"
)

func Get(id string) string {
	status, body, err := getAPI().Call("GET", fmt.Sprintf("/%s", id))

	if status == http.StatusOK {
		return body
	} else if status == http.StatusNotFound {
		return "Post not found."
	} else if status == http.StatusGone {
		return "Post unpublished."
	} else {
		return fmt.Sprintf("Problem getting post: %s. %v\n", status, err)
	}
}

func getAPI() *writeas.API {
	return writeas.GetAPI()
}
