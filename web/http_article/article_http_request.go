package http_article

type ArticlePostRequest struct {
	Title       string `validate:"required,min=1,max=250" json:"title"`
	Description string `validate:"required,min=1,max=2500" json:"description"`
	Image       string `json:"image"`
}

type ArticleUpdateRequest struct {
	ID          int    `validate:"required" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
