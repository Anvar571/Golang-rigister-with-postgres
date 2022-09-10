package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
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
	fmt.Println(data)
	_, e := db.Exec(query, data[0], data[1], data[2], data[3], data[4])
	if e != nil {
		panic(e)
	}
}

func (r *Rigister) delete(f_name string) {
	
}

func (r *Rigister) show() {
	var (
		f_name string
		lastname string
		job string
		parol int
		age int
	)
	readQuereData := `SELECT f_name, lastname, job, parol, age from rigister`
	row, err := db.Query(readQuereData)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&f_name, &lastname, &job, &parol, &age)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("no row ")
		case nil:
			fmt.Println(f_name, lastname, job, parol, age)
		default:
			fmt.Println("xatolik yuz berdi")
		}
	}
	
}

func main() {
	defer db.Close()
	Connect()
	var user = Rigister{}
	for {
		fmt.Println("[1]- Register\n[2]- Show data")
		var son int
		fmt.Scan(&son)
		switch son {
		case 1:
			var (
				f_name   string
				lastname string
				job      string
				parol    int
				age      int
			)
	
			fmt.Println("Enter first name")
			fmt.Scan(&f_name)
			fmt.Println("Enter last name")
			fmt.Scanf("%s", &lastname)
			fmt.Printf("Enter job name")
			fmt.Scanf("%s", &job)
			fmt.Println("Enter parol name")
			fmt.Scanln(&parol)
			fmt.Println("Enter age name")
			fmt.Scanln(&age)
	
			user.rigister(f_name, lastname, job, parol, age)
		case 2:
			user.show()
		}
	}
}