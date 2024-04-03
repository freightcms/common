package main

// Contact represents a contact that can be used to reach a company, organization, or individual.
type Contact struct {
	Id    string  // The unique identifier for the contact.
	Name  string  // The name of the contact.
	Email string  // The email address of the contact.
	Phone string  // The phone number of the contact.
	Fax   *string // The fax number of the contact.
}
