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

type Match struct {
	Id       int
	MatchId  int
	Accuracy int
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

	//Create a string matrix to output as a csv
	result := make([][]string, 0)

	//Adds the Header
	result = append(result, []string{"ContactID Source", "ContactID Match", "Accuracy"})

	for i, c := range contacts {

		//Since the range operator creates a copy of the values in the slice we do this to modify the original value
		contact := &contacts[i]
		contact.Checked = true

		matchs := findDuplicates(c, contacts)

		for _, m := range matchs {
			result = append(result, []string{strconv.Itoa(m.Id), strconv.Itoa(m.MatchId), strconv.Itoa(m.Accuracy)})
		}
	}

	output, err := os.Create("output.csv")

	if err != nil {
		log.Fatal("Error while creating the output file", err)
	}
	defer output.Close()

	writer := csv.NewWriter(output)
	defer writer.Flush()

	writer.WriteAll(result)
}

func findDuplicates(contact Contact, contactList []Contact) []Match {
	matchList := make([]Match, 0)
	for i, c := range contactList {
		if c.Checked {
			continue
		}
		accuracy := calculateDuplicate(contact, c)
		if accuracy > 0 {
			//Since the range operator creates a copy of the values in the slice we do this to modify the original value
			contactList[i].Checked = true
			m := Match{
				Id:       contact.Id,
				MatchId:  c.Id,
				Accuracy: accuracy,
			}

			matchList = append(matchList, m)
		}
	}
	return matchList
}

func calculateDuplicate(c1 Contact, c2 Contact) int {
	//returns a value from 0 to 5 depending on how likely the 2 contacts represent the same person
	accuracy := 0

	//If the emails are the same there are high chances that the contacts represent the same person
	if c1.Email == c2.Email {
		accuracy += 3
	}

	//First we compare if the firstnames are the same, if not, we do a letter by letter comparison
	//We do the same for the lastname
	if c1.FirstName == c2.FirstName {
		accuracy += 2
	} else {
		accuracy += compareStr(c1.FirstName, c2.LastName)
	}

	if c1.LastName == c2.LastName {
		accuracy += 2
	} else {
		accuracy += compareStr(c1.FirstName, c2.LastName)
	}

	//We cap the accuracy at 5 because of design choice
	if accuracy > 5 {
		accuracy = 5
	}
	return accuracy
}

func compareStr(str1 string, str2 string) int {
	n1 := strings.Split(str1, "")
	n2 := strings.Split(str2, "")

	var minLength int
	if len(n1) < len(n2) {
		minLength = len(n1)
	} else {
		minLength = len(n2)
	}

	for i := 0; i < minLength; i++ {
		if n1[i] != n2[i] {
			break
		}

		//if is the last iteration and the names are no different at this point there might be a coincidence, so the probability of being the same contact goes up
		if i == minLength-1 {
			return 1
		}
	}

	return 0
}

// Maps the data from the csv file into a list of Contacts
func getContactList(data [][]string) []Contact {

	contacts := make([]Contact, 1) //Initialize a list of contacts

	for _, contact := range data {
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
