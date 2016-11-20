package Model

import "time"

type Article struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	UpdateAt time.Time `json:"updateAt"`
	Author   string `json:"author"`
	Content  string `json:"content"`
}

type Articles []Article