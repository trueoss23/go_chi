package main

import (
	"database/sql"
	"log"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
	"github.com/trueoss23/go_chi/books/handlers"
	"github.com/trueoss23/go_chi/books/repo"
	"github.com/trueoss23/go_chi/books/usecases"
	cfg "github.com/trueoss23/go_chi/config"
	"github.com/trueoss23/go_chi/server"
)

func main() {
	dsn := cfg.Cfg.DbUser + ":" + cfg.Cfg.DbPass + "@tcp(localhost:3306)/" + cfg.Cfg.DbName
	Conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("No Conn!!", err)
	}
	defer Conn.Close()
	rep := repo.NewMySQLRepo(Conn)
	usecase := usecases.NewBookUseCase(rep)
	h := handlers.NewHandler(usecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/books", h.GetAllBooks)
	r.Post("/book", h.CreateBook)
	r.Delete("/book/{id}", h.DeleteBook)
	r.Get("/book/{id}", h.GetBook)
	server.StartServer(cfg.Cfg.AppPort, r)
}
