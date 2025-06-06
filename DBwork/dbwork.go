package dbwork

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type dbContact struct {
	id       int
	name     string
	phone    string
	birthday time.Time
	email    string
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "Vas9149970001"
	dbname   = "ContactManager"
)

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, _ = sql.Open("postgres", psqlconn)
}

func Disconnect() {
	db.Close()
}

func PrintDataBase() {
	rows, _ := db.Query("select * from contacts;")

	defer rows.Close()

	contacts := []dbContact{}

	for rows.Next() {
		var c dbContact

		rows.Scan(&c.id, &c.name, &c.phone, &c.birthday, &c.email)

		contacts = append(contacts, c)
	}

	for _, c := range contacts {
		fmt.Println(c.id, c.name, c.phone, c.birthday.Format(time.DateOnly), c.email)
	}
}

func AddContact(id int, name string, phone string, date time.Time, email string) {
	db.Exec("insert into contacts values ($1, $2, $3, $4, $5)", id, name, phone, date, email)
}

func DeleteContact(phone string) {
	row := db.QueryRow("select id from contacts where phone = $1", phone)
	var id int
	row.Scan(&id)

	db.Exec("delete from contacts where id = $1", id)
}
