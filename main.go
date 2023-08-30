package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/trueoss23/go_chi/models"
	"github.com/trueoss23/go_chi/db"
)

// InMemoryBookService is an implementation of BookService using in-memory storage
type InMemoryBookService struct {
	books []models.Book
}

func (s *InMemoryBookService) GetBooks() ([]models.Book, error) {
	return s.books, nil
}

func (s *InMemoryBookService) GetBookByID(id string) (*models.Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil // Return nil if book is not found
}

func (s *InMemoryBookService) CreateBook(book models.Book) error {
	s.books = append(s.books, book)
	return nil
}

func (s *InMemoryBookService) DeleteBook(id string) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return nil // Return nil if book is not found
}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	bookService := &InMemoryBookService{} // Create an instance of your BookService implementation

	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		books, err := bookService.GetBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(books)
	})

	r.Get("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		book, err := bookService.GetBookByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if book == nil {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(book)
	})

	// curl -X POST -H "Content-Type: application/json" -d '{
	// 	"id": "1",
	// 	"title": "The Great Gatsby",
	// 	"author": "F. Scott Fitzgerald"
	// }' http://localhost:3000/books

	r.Post("/books", func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = bookService.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Delete("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := bookService.DeleteBook(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":3000", r)
}
