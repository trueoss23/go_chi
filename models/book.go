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

type BookService interface {
	GetBooks() ([]Book, error)
	GetBookByID(id string) (*Book, error)
	CreateBook(book BookModel) error
	DeleteBook(id string) error
}
