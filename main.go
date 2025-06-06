package main

import (
	dbwork "ContactManager/DBwork"
	"fmt"
	"time"
)

func main() {

	dbwork.Connect()

	defer dbwork.Disconnect()

	dbwork.PrintDataBase()
	fmt.Println()

	dbwork.AddContact(4, "Ivan", "+79157771234", time.Date(1995, time.August, 15, 0, 0, 0, 0, time.UTC), "ivan.petrov@yandex.ru")
	dbwork.AddContact(55, "Alex", "8800", time.Date(1992, time.January, 28, 0, 0, 0, 0, time.UTC), "alex_brown@mail.ru")

	dbwork.DeleteContact("+79149970001")

	dbwork.PrintDataBase()
}
