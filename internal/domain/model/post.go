package model

type ID int64
type UserID int64

type Post struct {
	ID     ID     `json:"id" json:"id"`
	UserID UserID `json:"user_id" json:"user_id"`
	Title  string `json:"title" json:"title"`
	Body   string `json:"body" json:"body"`
}
