package main

import (
	manager "ContactManager/C_Manager"
	contact "ContactManager/Contact"
	"time"
)

func main() {

	cont1 := contact.NewContact("", "+79149970001", time.Date(2003, time.May, 22, 0, 0, 0, 0, time.UTC), "kamvasilii@gmail.com")
	cont2 := contact.NewContact("Ivan", "+79157771234", time.Date(1995, time.August, 15, 0, 0, 0, 0, time.UTC), "ivan.petrov@yandex.ru")
	cont3 := contact.NewContact("Maria", "+79162229876", time.Date(1988, time.November, 3, 0, 0, 0, 0, time.UTC), "m.smirnova@gmail.com")
	cont4 := contact.NewContact("Alexey", "+79153334455", time.Date(1992, time.January, 28, 0, 0, 0, 0, time.UTC), "alex_brown@mail.ru")
	cont5 := contact.NewContact("Elena", "+79158887766", time.Date(1998, time.July, 10, 0, 0, 0, 0, time.UTC), "elena_golubeva@outlook.com")

	cList := manager.Manager{}
	cList.AddContact(cont1)
	cList.AddContact(cont2)
	cList.AddContact(cont3)
	cList.AddContact(cont4)
	cList.AddContact(cont5)

	cList.DeleteContact("+79153334455")

	cList.PrinContacts()

}
