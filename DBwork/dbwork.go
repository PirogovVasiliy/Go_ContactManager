package dbwork

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type DBContact struct {
	id       int
	name     string
	phone    string
	birthday time.Time
	email    string
}

func NewDBContact(id int, name string, phone string, birthday time.Time, email string) DBContact {
	return DBContact{id: id, name: name, phone: phone, birthday: birthday, email: email}
}

func (contact *DBContact) GetPhone() string {
	return contact.phone
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

	contacts := []DBContact{}

	for rows.Next() {
		var c DBContact

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

func AddWithGo(contacts []DBContact) {
	var wg sync.WaitGroup
	for _, contact := range contacts {
		wg.Add(1)
		go func(c DBContact) {
			defer wg.Done()
			db.Exec("insert into contacts values ($1, $2, $3, $4, $5)", c.id, c.name, c.phone, c.birthday, c.email)
			time.Sleep(100 * time.Millisecond)
		}(contact)
	}
	wg.Wait()
}

func DeleteWithGo(contacts []DBContact) {
	var wg sync.WaitGroup
	for _, contact := range contacts {
		wg.Add(1)
		go func(c DBContact) {
			defer wg.Done()
			DeleteContact(c.GetPhone())
			time.Sleep(100 * time.Millisecond)
		}(contact)
	}
	wg.Wait()
}

func AddWithoutGo(contacts []DBContact) {
	for _, contact := range contacts {
		db.Exec("insert into contacts values ($1, $2, $3, $4, $5)", contact.id, contact.name, contact.phone, contact.birthday, contact.email)
		time.Sleep(100 * time.Millisecond)
	}
}

func DeleteWithoutGo(contacts []DBContact) {
	for _, contact := range contacts {
		func(c DBContact) {
			DeleteContact(c.GetPhone())
			time.Sleep(100 * time.Millisecond)
		}(contact)

	}
}

func GetContactsByName(name string) []int {
	res := make([]int, 0)

	rows, _ := db.Query("select id from contacts where name = $1", name)
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		res = append(res, id)
		time.Sleep(time.Second)
	}
	return res
}

func CountContacts() int {
	row := db.QueryRow("select count(*) from contacts")
	var count int
	row.Scan(&count)
	time.Sleep(time.Second)
	return count
}

func GetAllContacts() []DBContact {
	rows, _ := db.Query("select * from contacts;")
	defer rows.Close()

	contacts := []DBContact{}

	for rows.Next() {
		var c DBContact
		rows.Scan(&c.id, &c.name, &c.phone, &c.birthday, &c.email)
		contacts = append(contacts, c)
		time.Sleep(time.Second)
	}
	return contacts
}
