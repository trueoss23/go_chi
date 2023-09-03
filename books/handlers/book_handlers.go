package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/trueoss23/go_chi/books/usecases"
	"github.com/trueoss23/go_chi/domain/models"
)

type Handler struct {
	ctx     context.Context
	Usecase usecases.Usecase
}

func NewHandler(ctx context.Context, usecase usecases.Usecase) *Handler {
	return &Handler{
		ctx:     ctx,
		Usecase: usecase,
	}
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Usecase.GetAll()

	if err != nil {
		_, cancel := context.WithCancel(h.ctx)
		cancel()
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
		_, cancel := context.WithCancel(h.ctx)
		cancel()
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(book)
	bookInsert, err := h.Usecase.Insert(book)
	fmt.Println(bookInsert)

	if err != nil {
		_, cancel := context.WithCancel(h.ctx)
		cancel()
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data inserted successfully"))
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	err := h.Usecase.Delete(bookID)

	if err != nil {
		_, cancel := context.WithCancel(h.ctx)
		cancel()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Book deleted successfully")
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	book, err := h.Usecase.Get(bookID)

	if err != nil {
		_, cancel := context.WithCancel(h.ctx)
		cancel()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
