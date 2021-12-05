package base

type PostBody struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type FullPostBody struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Source  string `json:"source"`
	Picture string `json:"picture"`
}

type NewPostBody struct {
	Name       string `json:"name"`
	Content    string `json:"content"`
	Source     string `json:"source"`
	Picture    string `json:"picture"`
	CategoryID int    `json:"categoryId"`
}

type UpdatePostBody struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Source  string `json:"source"`
	Picture string `json:"picture"`
}
