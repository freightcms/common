package models

// HandlingUnitModel represents a handling unit model
type HandlingUnit struct {
	Id              string       `json:"id" bson:"id"`
	Description     string       `json:"description" bson:"description"`
	Weight          *Weight      `json:"weight" bson:"weight"`
	Dimensions      *Dimensions  `json:"dimensions" bson:"dimensions"`
	Volume          *Volume      `json:"volume" bson:"volume"`
	NFMC            string       `json:"nfmc" bson:"nfmc"`
	HarmonizedCode  string       `json:"harmonizedCode" bson:"harmonizedCode"`
	CountryOfOrigin string       `json:"countryOfOrigin" bson:"countryOfOrigin"`
	DeclaredValue   *Currency    `json:"declaredValue" bson:"declaredValue"`
	IsHazardous     bool         `json:"isHazardous" bson:"isHazardous"`
	IsStackable     bool         `json:"isStackable" bson:"isStackable"`
	Metadata        Metadata     `json:"metadata" bson:"metadata"`
	MinTemperature  *Temperature `json:"minTemperature" bson:"minTemperature"`
	MaxTemperature  *Temperature `json:"maxTemperature" bson:"maxTemperature"`
	References      []Reference  `json:"references" bson:"references"`
	Commodities     []Commodity  `json:"commodities" bson:"commodities"`
}
