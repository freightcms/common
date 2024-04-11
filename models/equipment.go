package models

type EquipmentType string

const (
	DryVan      EquipmentType = "DryVan"
	Reefer      EquipmentType = "Reefer"
	Flatbed     EquipmentType = "Flatbed"
	StepDeck    EquipmentType = "StepDeck"
	DoubleDrop  EquipmentType = "DoubleDrop"
	Lowboy      EquipmentType = "Lowboy"
	PowerOnly   EquipmentType = "PowerOnly"
	Container   EquipmentType = "Container"
	Chassis     EquipmentType = "Chassis"
	Tanker      EquipmentType = "Tanker"
	Hopper      EquipmentType = "Hopper"
	Pneumatic   EquipmentType = "Pneumatic"
	Dump        EquipmentType = "Dump"
	Livestock   EquipmentType = "Livestock"
	AutoCarrier EquipmentType = "AutoCarrier"
	Other       EquipmentType = "Other"
)

// EquipmentTypes is a list of all available equipment types.
var EquipmentTypes = []EquipmentType{
	DryVan,
	Reefer,
	Flatbed,
	StepDeck,
	DoubleDrop,
	Lowboy,
	PowerOnly,
	Container,
	Chassis,
	Tanker,
	Hopper,
	Pneumatic,
	Dump,
	Livestock,
	AutoCarrier,
	Other,
}
