package database

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	 connectStr := os.Getenv("CONNECTSTRING")
	 db, err := sql.Open("postgres", connectStr)

	 if err != nil {
		 log.Fatal("Error de conección: ", err)
	 }

	 return db
}