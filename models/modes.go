package models

type TransportationMode int

// TransportationMode represents the mode of transportation used to move goods from one location to another.
// See: https://kb.freightcms.com/modes/
const (
	Air TransportationMode = iota
	Ocean
	Rail
	LTL
	FTL
	VLTL
	Multimodal
	Intermodal
	Parcel
	Drayage
	RollOnRollOff
	LoadOnLoadOff
	BreakBulk
)
