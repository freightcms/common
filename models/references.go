package models

// ReferenceModel represents a reference to a model that can be used to identify a company, organization, or individual.
type ReferenceModel struct {
	Id          string  // Unique identifier for the reference for internal usage only.
	NameOrCode  string  // Name or code to relay to the consumer of what the reference may be for.
	Description *string // Description of the reference that tells more about the name or code that is implied.
	Value       string  // Any data that can be used by the end user to find information. JSON, XML, URLs, etc.
}
