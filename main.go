package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rigister"
	database = "rigister"
	password = "1"
)

var (
	dbSql   = fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s", host, port, user, password, database)
	db, err = sql.Open("postgres", dbSql)
)

func Connect() {
	if err != nil {
		log.Fatal(err)
	}
	error := db.Ping()
	if error != nil {
		log.Fatal(error)
	}
}

type Rigister struct {
	f_name   string
	lastname string
	job      string
	parol    int
	age      int
}

func (r *Rigister) rigister(f_name, lastname, job string, parol, age int) {
	data := strings.Split((fmt.Sprintf("%s %s %s %d %d", f_name, lastname, job, parol, age)), " ")
	query := `INSERT INTO rigister(f_name, lastname, job, parol, age)VALUES($1, $2, $3, $4, $5)`
	_, e := db.Exec(query, data[0], data[1], data[2], data[3], data[4])
	if e != nil {
		panic(e)
	}
}

func main() {
	defer db.Close()
	Connect()
	var user = Rigister{}
	var (
		f_name string
		lastname string
		job string
		parol int
		age int
	)

	fmt.Println("Enter first name")
	fmt.Scan(&f_name)
	fmt.Println("Enter last name")
	fmt.Scan(&lastname)
	fmt.Println("Enter job name")
	fmt.Scan(&job)
	fmt.Println("Enter parol name")
	fmt.Scan(&parol)
	fmt.Println("Enter age name")
	fmt.Scan(&age)
	
	user.rigister(f_name, lastname, job, parol, age)
}

// f_name, lastname, job string, parol, age int