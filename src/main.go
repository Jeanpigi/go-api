package main

import (
	"fmt"

	"github.com/Jeanpigi/go-api/src/database"
	"github.com/Jeanpigi/go-api/src/loadenv"
)

func main() {
	loadenv.LoadEnv()
	db := database.GetConnection()
	 fmt.Println(db)
}

