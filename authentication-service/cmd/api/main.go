package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/RianNegreiros/go-microservices/authentication-service/data"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting the authetication service")

	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
