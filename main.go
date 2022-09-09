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

type EmployeS interface {
	Connect()
	Rigister()
	Login()
	
}

type Rigister struct {
	f_name string
	lastname string
	parol int
	age int
	job string
}

func (r *Rigister) rigister(name, lastname, job string, parol, age int) string {
	query := `INSERT INTO rigister(f_name, lastname, parol, age, job)VALUES
	('anvar', 'abdurashidov', '123', '21', 'developer')`
	_, e := db.Exec(query)
	if e != nil {
		panic(e)
	}
}

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

	query := `INSERT INTO rigister(f_name, lastname, parol, age, job)VALUES
	('anvar', 'abdurashidov', '123', '21', 'developer')`
	_, e := db.Exec(query)
	if e != nil {
		panic(e)
	}
}