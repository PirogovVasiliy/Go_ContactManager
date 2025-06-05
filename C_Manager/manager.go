package manager

import (
	contact "ContactManager/Contact"
)

type Manager struct {
	contactList []contact.Contact
}

func NewManager(cl []contact.Contact) *Manager {
	return &Manager{contactList: cl}
}

func (cl *Manager) AddContact(con contact.Contact) {
	if (con == contact.Contact{}) {
		return
	} else {
		cl.contactList = append(cl.contactList, con)
	}
}

func (cl *Manager) DeleteContact(number string) {
	for i := 0; i < len(cl.contactList); i++ {
		if number == cl.contactList[i].GetPhone() {
			delSlice(&cl.contactList, i)
		}
	}
}

func (cl *Manager) PrinContacts() {
	for i := 0; i < len(cl.contactList); i++ {
		cl.contactList[i].Print()
	}
}

func (cl *Manager) GetContacts() []contact.Contact { return cl.contactList }

func delSlice(slice *[]contact.Contact, index int) {
	(*slice)[index] = (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
}
