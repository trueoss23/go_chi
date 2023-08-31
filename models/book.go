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

type BooksRepo interface {
	Connect() error
	Close() error
	GetAll() ([]Book, error)
	Get(id string) (Book, error)
	Insert(book BookModel) error
	Delete(id string) error
}
