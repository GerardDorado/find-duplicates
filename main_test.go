package main

import "testing"

func TestCompareStr_Similar(t *testing.T) {
	result := compareStr("Similar", "Simil")
	if result != 1 {
		t.Error("Failed comparing similar strings")
	}
}

func TestCompareStr_Different(t *testing.T) {
	result := compareStr("Similar", "Different")
	if result != 0 {
		t.Error("Failed comparing different strings")
	}
}

func TestCalculateDuplicate_Same_Contact(t *testing.T) {
	contact := Contact{FirstName: "John", LastName: "Doe", Email: "j.d@test.mail"}

	if calculateDuplicate(contact, contact) < 5 {
		t.Error("Failed identifying the same contact")
	}
}

func TestCalculateDuplicate_Different_Contact(t *testing.T) {
	contact_1 := Contact{FirstName: "John", LastName: "Doe", Email: "j.d@test.mail"}
	contact_2 := Contact{FirstName: "Joe", LastName: "Dee", Email: "j.d@other.mail"}

	if calculateDuplicate(contact_1, contact_2) > 1 {
		t.Error("Find similarity in different contact")
	}
}
