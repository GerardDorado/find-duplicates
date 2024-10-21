package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Zip       string
	Address   string
	//Add a boolean to mark that a contact was already checked for duplicates
	//Starts at false and is set to true if a contact is added to a list of possible duplicates
	Checked bool
}

func main() {
	file, err := os.Open("contacts.csv")
	if err != nil {
		log.Fatal("Error while opening the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error while reading the data from the file")
	}

	contacts := getContactList(data)

	for _, c := range contacts {
		log.Println(c.Address)
	}
}

// Maps the data from the csv file into a list of Contacts
func getContactList(data [][]string) []Contact {

	contacts := make([]Contact, 1) //Initialize a list of contacts

	for _, contact := range data {
		log.Println(contact)
		id, err := strconv.Atoi(contact[0])
		if err != nil {

			//since the first line of the csv is a reference for the colums is safe to ignore this error
			if contact[0] == "contactID" {
				continue
			}

			log.Fatal("Error while reading the id for contact", err)
		}

		c := Contact{
			Id:        id,
			FirstName: normalizeValue(contact[1]),
			LastName:  normalizeValue(contact[2]),
			Email:     normalizeValue(contact[3]),
			Zip:       normalizeValue(contact[4]),
			Address:   normalizeValue(contact[5]),
			Checked:   false,
		}
		contacts = append(contacts, c)
	}

	return contacts
}

func normalizeValue(str string) string {
	//Remove whitespace
	str = strings.ReplaceAll(str, " ", "")

	//Remove quotes
	str = strings.ReplaceAll(str, "\"", "")
	//Set to lowercase
	str = strings.ToLower(str)

	return str
}
