package responses

type BlogResponse struct {
	ID      uint   `json:"id"`
	UserId  uint   `json:"userId" form:"userId"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}