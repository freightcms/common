package models

// ReferenceModel represents a reference to a model that can be used to identify a company, organization, or individual.
type Reference struct {
	Id          string `json:"id" bson:"id"`                   // Unique identifier for the reference for internal usage only.
	NameOrCode  string `json:"nameOrCode" bson:"nameOrCode"`   // Name or code to relay to the consumer of what the reference may be for.
	Description string `json:"description" bson:"description"` // Description of the reference that tells more about the name or code that is implied.
	Value       string `json:"value" bson:"value"`             // Any data that can be used by the end user to find information. JSON, XML, URLs, etc.
}
