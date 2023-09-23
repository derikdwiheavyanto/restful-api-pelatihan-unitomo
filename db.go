package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initDBConnection() *sqlx.DB {
	connSTR := os.Getenv("DB_DSN")
	db, err := sqlx.Open("postgres", connSTR)

	if err != nil {
		log.Panic("couldn't connect to database", err)
	}
	fmt.Println("db successfully connected !")
	return db

}
