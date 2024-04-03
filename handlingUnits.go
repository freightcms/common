package models

// HandlingUnitModel represents a handling unit model
type HandlingUnitModel struct {
	Id              string
	Description     string
	Weight          *WeightModel
	Dimensions      *DimensionsModel
	Volume          *VolumeModel
	NFMC            *string
	HarmonizedCode  *string
	CountryOfOrigin *string
	DeclaredValue   *CurrencyModel
	IsHazardous     bool
	IsStackable     bool
	Metadata        Metadata
	MinTemperature  *TemperatureModel
	MaxTemperature  *TemperatureModel
	References      []ReferenceModel
	Commodities     []CommodityModel
}
