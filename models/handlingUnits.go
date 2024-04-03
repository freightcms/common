package models

// HandlingUnitModel represents a handling unit model
type HandlingUni struct {
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
	Commodities     []Commodity
}
