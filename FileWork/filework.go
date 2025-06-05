package filework

import (
	contact "ContactManager/Contact"
	"bufio"
	"os"
	"strings"
)

func LoadContacts() []contact.Contact {

	file, err := os.Open("contacts.txt")
	if err != nil {
		return []contact.Contact{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	loadedContacts := []contact.Contact{}
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		parsedContact := contact.ParseContactLine(trimmedLine)

		if (parsedContact != contact.Contact{}) {
			loadedContacts = append(loadedContacts, parsedContact)
		}
	}

	return loadedContacts
}

func SaveContacts(contacts []contact.Contact) {

	var linesToSave []string

	for i := 0; i < len(contacts); i++ {
		contactLine := contacts[i].FormatContactLine()
		linesToSave = append(linesToSave, contactLine)
	}

	data := strings.Join(linesToSave, "\n")

	os.WriteFile("contacts.txt", []byte(data), 0644)

}
