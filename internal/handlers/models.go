package handlers

type ArticleRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type ArticleResponse struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}
