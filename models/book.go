package models

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BookService provides methods for managing books
type BookService interface {
	GetBooks() ([]Book, error)
	GetBookByID(id string) (*Book, error)
	CreateBook(book Book) error
	DeleteBook(id string) error
}
