package jsonschemes

type Item struct {
	Slug    string `json:"slug"`
	Message string `json:"message"`
}

type Items []Item
