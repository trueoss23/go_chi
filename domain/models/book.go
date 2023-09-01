package models

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookModel struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}


