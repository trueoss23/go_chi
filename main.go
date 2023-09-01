package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	cfg "github.com/trueoss23/go_chi/config"
	"github.com/trueoss23/go_chi/handlers"
	"github.com/trueoss23/go_chi/books/usecases"
	"github.com/trueoss23/go_chi/books/repo"
)

func main() {
	mysqlrep = 
	rep = repo.NewBooksRepo(mysqlrep)
	usecase = usecases.NewBookUseCase(rep)
	db := *sql.DB
	h := &handlers.Handler{DB: db}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/books", h.GetAllBooks) // Используем методы хэндлера вместо анонимных функций
	r.Post("/book", h.CreateBook)
	r.Delete("/book/{id}", h.DeleteBook)
	r.Get("/book/{id}", h.GetBook)

	fmt.Println("Server listening on port " + cfg.Cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.Cfg.AppPort, r))
}
