package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RianNegreiros/go-microservices/authentication-service/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting the authetication service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Could not connect to the database")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DB_DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres is not ready yet. Retrying...")
			counts++
		} else {
			log.Println("Postgres is ready")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Waiting for 5 seconds...")
		time.Sleep(5 * time.Second)
		continue
	}
}
