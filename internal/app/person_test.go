package app

import (
	"testing"
)

func TestPersonCreate(t *testing.T) {
	var person Person = Person{}

	person.Id = "7da4941a-85a4-4c77-85bb-26b326c4105a"
	person.Email = "person@gmail.com"
	person.FirstName = "John"
	person.LastName = "Doe"
	person.Phone = "+1234567890"

	if person.Id != "7da4941a-85a4-4c77-85bb-26b326c4105a" {
		t.Errorf("Expected id to be '7da4941a-85a4-4c77-85bb-26b326c4105a', got '%s'", person.Id)
	}
	if person.Email != "person@gmail.com" {
		t.Errorf("Expected email to be 'person@gmail.com', got '%s'", person.Email)
	}
	if person.FirstName != "John" {
		t.Errorf("Expected FirstName to be 'John', got '%s'", person.FirstName)
	}
	if person.LastName != "Doe" {
		t.Errorf("Expected LastName to be 'Doe', got '%s'", person.LastName)
	}
	if person.Phone != "+1234567890" {
		t.Errorf("Expected Phone to be '+1234567890', got '%s'", person.Phone)
	}
}
