package dataobjects

// Source struct
type Source struct {
	ID   interface{} `json:"id"`
	Name string      `json:"name"`
}

//Article struct
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

//Articles Struct
type Articles struct {
	Articles []Article `json:"articles"`
}

//NewsResponse struct
type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
