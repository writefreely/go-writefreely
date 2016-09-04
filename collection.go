package writeas

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
