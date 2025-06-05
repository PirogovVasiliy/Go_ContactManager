package main

import (
	manager "ContactManager/C_Manager"
	contact "ContactManager/Contact"
	filework "ContactManager/FileWork"
	"time"
)

func main() {

	cList := manager.NewManager(filework.LoadContacts())

	cList.PrinContacts()

	cont1 := contact.NewContact("Vasya", "+79149970001", time.Date(2003, time.May, 22, 0, 0, 0, 0, time.UTC), "kamvasilii@gmail.com")
	cList.AddContact(cont1)

	cList.DeleteContact("+79158887766")

	filework.SaveContacts(cList.GetContacts())
}
