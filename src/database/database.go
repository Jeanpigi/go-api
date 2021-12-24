package database

import (
	"database/sql"
	"os"

	handlererror "github.com/Jeanpigi/go-api/src/handlerError"
	_ "github.com/lib/pq"
)


func GetConnection() *sql.DB {	
	connectStr := os.Getenv("CONNECTSTRING")
	db, err := sql.Open("postgres", connectStr)

	handlererror.CheckError(err, "")

	return db
}