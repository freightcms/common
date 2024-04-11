package models

type WeightUnit string

const (
	Kilogram WeightUnit = "kg"
	Gram     WeightUnit = "g"
	Pound    WeightUnit = "lb"
	Ounce    WeightUnit = "oz"
)

// WeightModel represents a model for weight
type Weight struct {
	Value float64    `json:"value" bson:"value"` // The weight value
	Unit  WeightUnit `json:"unit" bson:"unit"`   // The unit of the weight. Possible values are: "kg", "lb"
}
