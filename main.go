package main

import (
	"fmt"
	"log"
	"net/http"

	cfg "github.com/trueoss23/go_chi/config"
	"github.com/trueoss23/go_chi/db"
	"github.com/trueoss23/go_chi/routes"
)

func main() {
	db := &db.MySQLDatabase{}
	err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer db.Close()
	r := routes.SetupRoutes(db)

	fmt.Println(cfg.Cfg.AppPort, cfg.Cfg.AppHost, cfg.Cfg.DbUser, cfg.Cfg.DbPass, cfg.Cfg.DbName)
	fmt.Println("Server listening on port " + cfg.Cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.Cfg.AppPort, r))
}
