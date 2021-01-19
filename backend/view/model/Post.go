package model

// 投稿
type Post struct {
	// ID
	ID string `json:"id"`
	// タイトル
	Title string `json:"title"`
	// 投稿者
	AuthorID string `json:"authorId"`
}
