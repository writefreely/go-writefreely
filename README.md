# go-writeas

Official Write.as Go client library.

## Installation

```bash
go get github.com/writeas/go-writeas
```

## Documentation

See all functionality and usages in the [API documentation](https://writeas.github.io/docs/).

### Example usage

```go
import "github.com/writeas/go-writeas"

func main() {
	// Create the client
	c := writeas.NewClient()

	// Publish a post
	p, err := c.CreatePost(&PostParams{
		Title:   "Title!",
		Content: "This is a post.",
		Font:    "sans",
	})
	if err != nil {
		// Perhaps show err.Error()
	}

	// Update a published post
	p, err := c.UpdatePost(&PostParams{
		ID:      "3psnxyhqxy3hq",
		Token:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Content: "Now it's been updated!",
	})
	if err != nil {
		// handle
	}

	// Get a published post
	p, err := c.GetPost("3psnxyhqxy3hq")
	if err != nil {
		// handle
	}

	// Delete a post
	err := c.DeletePost(&PostParams{
		ID:    "3psnxyhqxy3hq",
		Token: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	})
}
```

## Contributing

The library covers our usage, but might not be comprehensive of the API. So we always welcome contributions and improvements from the community. Before sending pull requests, make sure you've done the following:

* Run `go fmt` on all updated .go files.
* Document all structs and funcs.

## License

MIT
