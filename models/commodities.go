package models

// Commodity represents a commodity model
type Commodity struct {
	Id              string
	Description     string
	Weight          *Weight
	Dimensions      *Dimensions
	Volume          *Volume
	NFMC            *string
	HarmonizedCode  *string
	CountryOfOrigin *string
	DeclaredValue   *Currency
	IsHazardous     bool
	IsStackable     bool
	Metadata        Metadata
	MinTemperature  *Temperature
	MaxTemperature  *Temperature
	References      []Reference
}
