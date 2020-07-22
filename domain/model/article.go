package model

type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type ArticleRepository interface {
	All() []Article
	FindByID(ID int) Article
}
