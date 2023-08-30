package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/trueoss23/go_chi/db"
	"github.com/trueoss23/go_chi/models"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	db := &db.MySQLDatabase{}
	err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer db.Close()

	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		books, err := db.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	})

	r.Post("/books", func(w http.ResponseWriter, r *http.Request) {
		var book models.BookModel
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		fmt.Println(book)
		err = db.Insert(book)
		if err != nil {
			http.Error(w, "Failed to insert data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Data inserted successfully"))
	})

	r.Delete("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "id")

		err := db.Delete(bookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Book deleted successfully")
	})

	r.Get("/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "id")

		book, err := db.Get(bookID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	})

	return r
}
