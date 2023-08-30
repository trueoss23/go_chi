package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/trueoss23/go_chi/db"
    "github.com/trueoss23/go_chi/routes"
)

func main() {
    r := routes.SetupRoutes()

    db := &db.MySQLDatabase{}
    err := db.Connect()
    if err != nil {
        log.Fatal("Error connecting to db:", err)
    }
    defer db.Close()

    fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
