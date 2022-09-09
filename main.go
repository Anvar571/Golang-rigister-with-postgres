package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rigister"
	database = "rigister"
	password = "1"
)

func main() {
	dbSql := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s", host, port, user, password, database)
	db, err := sql.Open("postgres", dbSql)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	error := db.Ping()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("connection success")
}
