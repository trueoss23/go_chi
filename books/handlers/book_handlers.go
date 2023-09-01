package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/trueoss23/go_chi/books/usecases"
	"github.com/trueoss23/go_chi/domain/models"
)

type Handler struct {
	DB usecases.Usecase
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.DB.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.BookModel
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(book)
	err = h.DB.Insert(book)

	if err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data inserted successfully"))
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	err := h.DB.Delete(bookID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Book deleted successfully")
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	book, err := h.DB.Get(bookID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
