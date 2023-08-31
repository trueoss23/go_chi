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
	r := routes.SetupRoutes(db)

	fmt.Println("Server listening on port " + cfg.Cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.Cfg.AppPort, r))
}
