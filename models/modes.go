package models

type TransportationMode string

// TransportationMode represents the mode of transportation used to move goods from one location to another.
// See: https://kb.freightcms.com/modes/
const (
	Air           TransportationMode = "Air"
	Ocean         TransportationMode = "Ocean"
	Rail          TransportationMode = "Rail"
	LTL           TransportationMode = "LTL"
	FTL           TransportationMode = "FTL"
	VLTL          TransportationMode = "VLTL"
	Multimodal    TransportationMode = "Multimodal"
	Intermodal    TransportationMode = "Intermodal"
	Parcel        TransportationMode = "Parcel"
	Drayage       TransportationMode = "Drayage"
	RollOnRollOff TransportationMode = "RoRo"
	LoadOnLoadOff TransportationMode = "LoLo"
	BreakBulk     TransportationMode = "BreakBulk"
)

// TransportationModes is a list of all available transportation modes.
var TransportationModes = []TransportationMode{
	Air,
	Ocean,
	Rail,
	LTL,
	FTL,
	VLTL,
	Multimodal,
	Intermodal,
	Parcel,
	Drayage,
	RollOnRollOff,
	LoadOnLoadOff,
	BreakBulk,
}
