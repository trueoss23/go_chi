package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/trueoss23/go_chi/books/repo"
	"github.com/trueoss23/go_chi/books/usecases"
	cfg "github.com/trueoss23/go_chi/config"
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
	startServer(cfg.Cfg.AppPort, server.R)

}
