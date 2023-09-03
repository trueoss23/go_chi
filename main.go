package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
	"github.com/trueoss23/go_chi/books/handlers"
	"github.com/trueoss23/go_chi/books/repo"
	"github.com/trueoss23/go_chi/books/usecases"
	cfg "github.com/trueoss23/go_chi/config"
)

func main() {
	dsn := cfg.Cfg.DbUser + ":" + cfg.Cfg.DbPass + "@tcp(localhost:3306)/" + cfg.Cfg.DbName
	Conn, err := sql.Open("mysql", dsn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal("No Conn!!", err)
	}
	defer Conn.Close()
	rep := repo.NewMySQLRepo(ctx, Conn)
	usecase := usecases.NewBookUseCase(rep)
	h := handlers.NewHandler(usecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/books", h.GetAllBooks)
	r.Post("/book", h.CreateBook)
	r.Delete("/book/{id}", h.DeleteBook)
	r.Get("/book/{id}", h.GetBook)
	StartServer(ctx, cfg.Cfg.AppPort, r)
}
