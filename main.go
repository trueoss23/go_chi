package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/trueoss23/go_chi/books/handlers"
	"github.com/trueoss23/go_chi/books/repo"
	"github.com/trueoss23/go_chi/books/usecases"
	cfg "github.com/trueoss23/go_chi/config"
)

func main() {
	dsn := cfg.Cfg.DbUser + ":" + cfg.Cfg.DbPass + "@tcp(localhost:3306)/" + cfg.Cfg.DbName
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("No conn!!", err)
	}
	defer conn.Close()
	rep := repo.NewBooksRepo(conn)
	usecase := usecases.NewBookUseCase(rep)
	// db := *sql.DB
	h := &handlers.Handler{DB: usecase}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/books", h.GetAllBooks)
	r.Post("/book", h.CreateBook)
	r.Delete("/book/{id}", h.DeleteBook)
	r.Get("/book/{id}", h.GetBook)

	fmt.Println("Server listening on port " + cfg.Cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.Cfg.AppPort, r))
}
