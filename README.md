# go-writeas

[![godoc](https://godoc.org/go.code.as/writeas.v1?status.svg)](https://godoc.org/go.code.as/writeas.v1)

Official Write.as Go client library.

## Installation

```bash
go get go.code.as/writeas.v1
```

## Documentation

See all functionality and usages in the [API documentation](https://developer.write.as/docs/api/).

### Example usage

```go
import "go.code.as/writeas.v1"

func main() {
	// Create the client
	c := writeas.NewClient()

	// Publish a post
	p, err := c.CreatePost(&writeas.PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		// Perhaps show err.Error()
	}

	// Save token for later, since it won't ever be returned again
	token := p.Token

	// Update a published post
	p, err = c.UpdatePost(&writeas.PostParams{
		OwnedPostParams: writeas.OwnedPostParams{
			ID:    p.ID,
			Token: token,
		},
		Content: "Now it's been updated!",
	})
	if err != nil {
		// handle
	}

	// Get a published post
	p, err = c.GetPost(p.ID)
	if err != nil {
		// handle
	}

	// Delete a post
	err = c.DeletePost(&writeas.PostParams{
		OwnedPostParams: writeas.OwnedPostParams{
			ID:    p.ID,
			Token: token,
		},
	})
}
```

## Contributing

The library covers our usage, but might not be comprehensive of the API. So we always welcome contributions and improvements from the community. Before sending pull requests, make sure you've done the following:

* Run `go fmt` on all updated .go files.
* Document all exported structs and funcs.

## License

MIT
