package contact

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Contact struct {
	name        string
	phoneNumber string
	birthday    time.Time
	email       string
}

func NewContact(name string, phone string, birthday time.Time, email string) Contact {
	if name == "" {
		return Contact{}
	}
	if phone == "" {
		return Contact{}
	}
	if email == "" {
		return Contact{}
	}

	return Contact{name: name, phoneNumber: phone, birthday: birthday, email: email}
}

func (c *Contact) GetName() string { return c.name }

func (c *Contact) GetPhone() string { return c.phoneNumber }

func (c *Contact) GetBirthday() time.Time { return c.birthday }

func (c *Contact) GetEmail() string { return c.email }

func (c *Contact) SetName(newName string) {
	if newName == "" {
		return
	} else {
		c.name = newName
	}
}

func (c *Contact) SetPhone(newPhone string) {
	if newPhone != "" {
		c.phoneNumber = newPhone
	} else {
		return
	}
}

func (c *Contact) SetBirthday(newBirthday time.Time) { c.birthday = newBirthday }

func (c *Contact) SetEmail(newEmail string) {
	if newEmail == "" {
		return
	} else {
		c.email = newEmail
	}
}

func (c *Contact) Print() {
	fmt.Println(color.RedString("Name:"), c.name)
	fmt.Println(color.BlueString("Phone Number:"), c.phoneNumber)
	fmt.Println(color.YellowString("Birthday:"), c.birthday.Format(time.DateOnly))
	fmt.Println(color.MagentaString("Email:"), c.email)
	fmt.Println()
}

func (c *Contact) FormatContactLine() string {
	return c.name + ";" + c.phoneNumber + ";" + c.birthday.Format(time.DateOnly) + ";" + c.email
}

func ParseContactLine(line string) Contact {

	parts := strings.Split(line, ";")

	name := strings.TrimSpace(parts[0])
	phone := strings.TrimSpace(parts[1])
	birthdayStr := strings.TrimSpace(parts[2])
	email := strings.TrimSpace(parts[3])

	birthday, err := time.Parse(time.DateOnly, birthdayStr)
	if err != nil {
		return Contact{}
	}

	if name == "" || phone == "" || email == "" {
		return Contact{}
	}

	return Contact{
		name:        name,
		phoneNumber: phone,
		birthday:    birthday,
		email:       email,
	}
}
