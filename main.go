package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

// Create - Insert
func (r *Rigister) rigister(f_name, lastname, job string, parol, age int) {
	data := strings.Split((fmt.Sprintf("%s %s %s %d %d", f_name, lastname, job, parol, age)), " ")
	query := `INSERT INTO rigister(f_name, lastname, job, parol, age)VALUES($1, $2, $3, $4, $5)`
	fmt.Println(data)
	_, e := db.Exec(query, data[0], data[1], data[2], data[3], data[4])
	if e != nil {
		panic(e)
	}
}
// update
func (r *Rigister) updateData(id string, f_name string, parol int) {
	updateQuery := `UPDATE rigister SET f_name=$2, parol=$3 where id = $1`
	_, err := db.Exec(updateQuery, id, f_name, parol)
	if err != nil {
		panic(err)
	}
}

// delete
func (r *Rigister) deleteData(f_name string) {
	if len(f_name) < 3 {
		fmt.Println("con not letters")
		os.Exit(1)
	}
	deleteQuery := `delete from rigister where f_name = $1`
	_, err := db.Exec(deleteQuery, f_name)
	if err != nil {
		panic(err)
	}
	fmt.Println("delete successfull")
}

// Read
func (r *Rigister) show() {
	var (
		f_name   string
		lastname string
		job      string
		parol    int
		age      int
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
		fmt.Println("[1] - Register\n[2] - Show data\n[3] - Delete user\n[4] - Update user")
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
		case 3:
			var name string
			fmt.Println("Enter delete name")
			fmt.Scan(&name)
			user.deleteData(name)
		case 4:
			var (
				f_name string
				parolQ int
				id string
			)
			fmt.Println("Qaysi ididagi userni yangilamoqchisiz: ")
			fmt.Scan(&id)
			fmt.Println("Enter new f_name")
			fmt.Scan(&f_name)
			fmt.Println("Enter new parol")
			fmt.Scan(&parolQ)
			user.updateData(id, f_name, parolQ)
		}
	}
}
