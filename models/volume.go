package models

// Unit of Volume type for measurement
// see https://en.wikipedia.org/wiki/Unit_of_volume
type VolumeUnit string

const (
	CubicMetre     VolumeUnit = "m³"
	Barrel         VolumeUnit = "bbl"
	CubicFoot      VolumeUnit = "ft³"
	Litre          VolumeUnit = "L"
	CubicDecimetre VolumeUnit = "dm³"
	Gallon         VolumeUnit = "gal"
	Pint           VolumeUnit = "pt"
	CubicInch      VolumeUnit = "in³"
	CubicYard      VolumeUnit = "yd³"
	Cord           VolumeUnit = "cord"
	BoardFoot      VolumeUnit = "fbm"
)

type Volume struct {
	Value float64    `json:"value" bson:"value"` // The volume value
	Unit  VolumeUnit `json:"unit" bson:"unit"`   // The unit of the volume
}
