package models

// Commodity represents a commodity model
// commodities represent the goods being shipped
// see https://kb.freightcms.com/packaging/commodities/
type Commodity struct {
	Id              string       `json:"id" bson:"id"`                   // unique identifier for a commodity
	Description     string       `json:"description" bson:"description"` // (Optional) description of a commodity
	Weight          *Weight      `json:"weight" bson:"weight"`           // weight of a commodity
	Dimensions      *Dimensions  `json:"dimensions" bson:"dimensions"`
	Volume          *Volume      `json:"volume" bson:"volume"`                   // volume of a commodity
	NFMC            string       `json:"nfmc" bson:"nmfc"`                       // National Motor Freight Classification
	CountryOfOrigin string       `json:"countryOfOrigin" bson:"countryOfOrigin"` // country of origin of a commodity
	DeclaredValue   *Currency    `json:"declaredValue" bson:"declaredValue"`     // declared value of a commodity
	IsHazardous     bool         `json:"isHazardous" bson:"isHazardous"`         // whether the commodity is hazardous
	IsStackable     bool         `json:"isStackable" bson:"isStackable"`         // whether the commodity can be stacked
	Metadata        Metadata     `json:"metadata" bson:"metadata"`               // metadata information
	MinTemperature  *Temperature `json:"minTemperature" bson:"minTemperature"`   // maximum temperature required for a commodity
	MaxTemperature  *Temperature `json:"maxTemperature" bson:"maxTemperature"`   // minimum temperature required for a commodity
	References      []Reference  `json:"references" bson:"references"`           // any customer references to a commodity
	UOM             string       `json:"uom" bson:"uom"`                         // unit of measure
}
