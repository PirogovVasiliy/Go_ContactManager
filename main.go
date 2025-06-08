package main

import (
	dbwork "ContactManager/DBwork"
	"fmt"
	"sync"
	"time"
)

func main() {

	dbwork.Connect()

	defer dbwork.Disconnect()

	//contacts := createContacts()

	//start := time.Now()
	//dbwork.AddWithGo(contacts)
	//dbwork.AddWithoutGo(contacts)
	//dbwork.DeleteWithGo(contacts) //10.0923154
	//dbwork.DeleteWithoutGo(contacts)
	//fmt.Println(time.Since(start).Seconds())

	//withGO()
	//withoutGO()
}

func withGO() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println(len(dbwork.GetAllContacts()))
	}()
	go func() {
		defer wg.Done()
		fmt.Println(len(dbwork.GetContactsByName("Ivan")))
	}()
	go func() {
		defer wg.Done()
		fmt.Println(dbwork.CountContacts())
	}()
	wg.Wait()
	fmt.Println(time.Since(start).Seconds())
}

func withoutGO() {
	start := time.Now()
	fmt.Println(len(dbwork.GetAllContacts()))
	fmt.Println(len(dbwork.GetContactsByName("Ivan")))
	fmt.Println(dbwork.CountContacts())
	fmt.Println(time.Since(start).Seconds())
}

func createContacts() []dbwork.DBContact {
	contacts := []dbwork.DBContact{}

	baseID := 100
	basePhonePrefix := "+999"
	baseEmailDomain := "@example.com"
	baseDate := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 100; i++ {
		id := baseID + i
		name := fmt.Sprintf("GeneratedUser_%d", i)
		phone := fmt.Sprintf("%s%04d", basePhonePrefix, i)
		email := fmt.Sprintf("user_%d%s", i, baseEmailDomain)
		birthday := baseDate.AddDate(0, 0, i)

		newContact := dbwork.NewDBContact(id, name, phone, birthday, email)

		contacts = append(contacts, newContact)
	}

	return contacts
}
