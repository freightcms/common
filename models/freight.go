package models

// FreightClass represents the freight class of a shipment.
// see https://kb.freightcms.com/glossery/freight-class/
type FreightClass int

const (
	Class50 FreightClass = iota + 50
	Class55
	Class60
	Class65
	Class70
	Class77_55
	Class85
	Class92_5
	Class100
	Class110
	Class125
	Class150
	Class175
	Class200
	Class250
	Class300
	Class400
	Class500
)

// FreightClasses is a list of all available freight classes.
var FreightClasses = []FreightClass{
	Class50,
	Class55,
	Class60,
	Class65,
	Class70,
	Class77_55,
	Class85,
	Class92_5,
	Class100,
	Class110,
	Class125,
	Class150,
	Class175,
	Class200,
	Class250,
	Class300,
	Class400,
	Class500,
}
