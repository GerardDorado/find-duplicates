package main

import (
	"encoding/csv"
	"log"
	"os"
)

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

	for _, contact := range data {
		log.Println(contact)
	}
}
