package models

// Contact represents a contact that can be used to reach a company, organization, or individual.
type Contact struct {
	Id    string `json:"id" bson:"id"`       // The unique identifier for the contact.
	Name  string `json:"name" bson:"name"`   // The name of the contact.
	Email string `json:"email" bson:"email"` // The email address of the contact.
	Phone string `json:"phone" bson:"phone"` // The phone number of the contact.
	Fax   string `json:"fax" bson:"fax"`     // The fax number of the contact.
}
