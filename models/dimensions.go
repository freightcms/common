package models

type DimensionUnit string

const (
	Millimeter DimensionUnit = "mm"
	Centimeter DimensionUnit = "cm"
	Meter      DimensionUnit = "m"
	Inch       DimensionUnit = "in"
	Foot       DimensionUnit = "ft"
	Yard       DimensionUnit = "yd"
	Mile       DimensionUnit = "mi"
	KiloMeter  DimensionUnit = "km"
)

var DimensionUnits = []DimensionUnit{
	Millimeter,
	Centimeter,
	Meter,
	Inch,
	Foot,
	Yard,
	Mile,
	KiloMeter,
}

// DimensionsModel represents a model for dimensions
type Dimensions struct {
	Length float64       `json:"length" bson:"length"` // The length dimension
	Width  float64       `json:"width" bson:"width"`   // The width dimension
	Height float64       `json:"height" bson:"height"` // The height dimension
	Unit   DimensionUnit `json:"unit" bson:"unit"`     // The unit of the dimensions
}

// ToMillimeter converts the dimensions to millimeters
func (d *Dimensions) ToMillimeter() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return d
	case Centimeter:
		return &Dimensions{
			Length: d.Length * 10,
			Width:  d.Width * 10,
			Height: d.Height * 10,
			Unit:   Millimeter,
		}
	case Meter:
		return &Dimensions{
			Length: d.Length * 1000,
			Width:  d.Width * 1000,
			Height: d.Height * 1000,
			Unit:   Millimeter,
		}
	case Inch:
		return &Dimensions{
			Length: d.Length * 25.4,
			Width:  d.Width * 25.4,
			Height: d.Height * 25.4,
			Unit:   Millimeter,
		}
	case Foot:
		return &Dimensions{
			Length: d.Length * 304.8,
			Width:  d.Width * 304.8,
			Height: d.Height * 304.8,
			Unit:   Millimeter,
		}
	case Yard:
		return &Dimensions{
			Length: d.Length * 914.4,
			Width:  d.Width * 914.4,
			Height: d.Height * 914.4,
			Unit:   Millimeter,
		}
	case Mile:
		return &Dimensions{
			Length: d.Length * 1609344,
			Width:  d.Width * 1609344,
			Height: d.Height * 1609344,
			Unit:   Millimeter,
		}
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 1000000,
			Width:  d.Width * 1000000,
			Height: d.Height * 1000000,
			Unit:   Millimeter,
		}
	default:
		return nil
	}
}

// ToCentimeter converts the dimensions to centimeters
func (d *Dimensions) ToCentimeter() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 10,
			Width:  d.Width / 10,
			Height: d.Height / 10,
			Unit:   Centimeter,
		}
	case Centimeter:
		return d
	case Meter:
		return &Dimensions{
			Length: d.Length * 100,
			Width:  d.Width * 100,
			Height: d.Height * 100,
			Unit:   Centimeter,
		}
	case Inch:
		return &Dimensions{
			Length: d.Length * 2.54,
			Width:  d.Width * 2.54,
			Height: d.Height * 2.54,
			Unit:   Centimeter,
		}
	case Foot:
		return &Dimensions{
			Length: d.Length * 30.48,
			Width:  d.Width * 30.48,
			Height: d.Height * 30.48,
			Unit:   Centimeter,
		}
	case Yard:
		return &Dimensions{
			Length: d.Length * 91.44,
			Width:  d.Width * 91.44,
			Height: d.Height * 91.44,
			Unit:   Centimeter,
		}
	case Mile:
		return &Dimensions{
			Length: d.Length * 160934.4,
			Width:  d.Width * 160934.4,
			Height: d.Height * 160934.4,
			Unit:   Centimeter,
		}
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 100000,
			Width:  d.Width * 100000,
			Height: d.Height * 100000,
			Unit:   Centimeter,
		}
	default:
		return nil
	}
}

// ToMeter converts the dimensions to meters
func (d *Dimensions) ToMeter() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 1000,
			Width:  d.Width / 1000,
			Height: d.Height / 1000,
			Unit:   Meter,
		}
	case Centimeter:
		return &Dimensions{
			Length: d.Length / 100,
			Width:  d.Width / 100,
			Height: d.Height / 100,
			Unit:   Meter,
		}
	case Meter:
		return d
	case Inch:
		return &Dimensions{
			Length: d.Length * 0.0254,
			Width:  d.Width * 0.0254,
			Height: d.Height * 0.0254,
			Unit:   Meter,
		}
	case Foot:
		return &Dimensions{
			Length: d.Length * 0.3048,
			Width:  d.Width * 0.3048,
			Height: d.Height * 0.3048,
			Unit:   Meter,
		}
	case Yard:
		return &Dimensions{
			Length: d.Length * 0.9144,
			Width:  d.Width * 0.9144,
			Height: d.Height * 0.9144,
			Unit:   Meter,
		}
	case Mile:
		return &Dimensions{
			Length: d.Length * 1609.344,
			Width:  d.Width * 1609.344,
			Height: d.Height * 1609.344,
			Unit:   Meter,
		}
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 1000,
			Width:  d.Width * 1000,
			Height: d.Height * 1000,
			Unit:   Meter,
		}
	default:
		return nil
	}
}

// ToInch converts the dimensions to inches
func (d *Dimensions) ToInch() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 25.4,
			Width:  d.Width / 25.4,
			Height: d.Height / 25.4,
			Unit:   Inch,
		}
	case Centimeter:
		return &Dimensions{
			Length: d.Length / 2.54,
			Width:  d.Width / 2.54,
			Height: d.Height / 2.54,
			Unit:   Inch,
		}
	case Meter:
		return &Dimensions{
			Length: d.Length / 0.0254,
			Width:  d.Width / 0.0254,
			Height: d.Height / 0.0254,
			Unit:   Inch,
		}
	case Inch:
		return d
	case Foot:
		return &Dimensions{
			Length: d.Length * 12,
			Width:  d.Width * 12,
			Height: d.Height * 12,
			Unit:   Inch,
		}
	case Yard:
		return &Dimensions{
			Length: d.Length * 36,
			Width:  d.Width * 36,
			Height: d.Height * 36,
			Unit:   Inch,
		}
	case Mile:
		return &Dimensions{
			Length: d.Length * 63360,
			Width:  d.Width * 63360,
			Height: d.Height * 63360,
			Unit:   Inch,
		}
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 39370.1,
			Width:  d.Width * 39370.1,
			Height: d.Height * 39370.1,
			Unit:   Inch,
		}
	default:
		return nil
	}
}

// ToFoot converts the dimensions to feet
func (d *Dimensions) ToFoot() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 304.8,
			Width:  d.Width / 304.8,
			Height: d.Height / 304.8,
			Unit:   Foot,
		}
	case Centimeter:
		return &Dimensions{
			Length: d.Length / 30.48,
			Width:  d.Width / 30.48,
			Height: d.Height / 30.48,
			Unit:   Foot,
		}
	case Meter:
		return &Dimensions{
			Length: d.Length / 0.3048,
			Width:  d.Width / 0.3048,
			Height: d.Height / 0.3048,
			Unit:   Foot,
		}
	case Inch:
		return &Dimensions{
			Length: d.Length / 12,
			Width:  d.Width / 12,
			Height: d.Height / 12,
			Unit:   Foot,
		}
	case Foot:
		return d
	case Yard:
		return &Dimensions{
			Length: d.Length * 3,
			Width:  d.Width * 3,
			Height: d.Height * 3,
			Unit:   Foot,
		}
	case Mile:
		return &Dimensions{
			Length: d.Length * 5280,
			Width:  d.Width * 5280,
			Height: d.Height * 5280,
			Unit:   Foot,
		}
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 3280.84,
			Width:  d.Width * 3280.84,
			Height: d.Height * 3280.84,
			Unit:   Foot,
		}
	default:
		return nil
	}
}

// ToYard converts the dimensions to yards
func (d *Dimensions) ToYard() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 914.4,
			Width:  d.Width / 914.4,
			Height: d.Height / 914.4,
			Unit:   Yard,
		}
	case Centimeter:
		return &Dimensions{
			Length: d.Length / 91.44,
			Width:  d.Width / 91.44,
			Height: d.Height / 91.44,
			Unit:   Yard,
		}
	case Meter:
		return &Dimensions{
			Length: d.Length / 0.9144,
			Width:  d.Width / 0.9144,
			Height: d.Height / 0.9144,
			Unit:   Yard,
		}
	case Inch:
		return &Dimensions{
			Length: d.Length / 36,
			Width:  d.Width / 36,
			Height: d.Height / 36,
			Unit:   Yard,
		}
	case Foot:
		return &Dimensions{
			Length: d.Length / 3,
			Width:  d.Width / 3,
			Height: d.Height / 3,
			Unit:   Yard,
		}
	case Yard:
		return d
	case Mile:
		return &Dimensions{
			Length: d.Length * 1760,
			Width:  d.Width * 1760,
			Height: d.Height * 1760,
			Unit:   Yard,
		}
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 1093.61,
			Width:  d.Width * 1093.61,
			Height: d.Height * 1093.61,
			Unit:   Yard,
		}
	default:
		return nil
	}
}

// ToMile converts the dimensions to miles
func (d *Dimensions) ToMile() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 1609344,
			Width:  d.Width / 1609344,
			Height: d.Height / 1609344,
			Unit:   Mile,
		}
	case Centimeter:
		return &Dimensions{
			Length: d.Length / 160934.4,
			Width:  d.Width / 160934.4,
			Height: d.Height / 160934.4,
			Unit:   Mile,
		}
	case Meter:
		return &Dimensions{
			Length: d.Length / 1609.344,
			Width:  d.Width / 1609.344,
			Height: d.Height / 1609.344,
			Unit:   Mile,
		}
	case Inch:
		return &Dimensions{
			Length: d.Length / 63360,
			Width:  d.Width / 63360,
			Height: d.Height / 63360,
			Unit:   Mile,
		}
	case Foot:
		return &Dimensions{
			Length: d.Length / 5280,
			Width:  d.Width / 5280,
			Height: d.Height / 5280,
			Unit:   Mile,
		}
	case Yard:
		return &Dimensions{
			Length: d.Length / 1760,
			Width:  d.Width / 1760,
			Height: d.Height / 1760,
			Unit:   Mile,
		}
	case Mile:
		return d
	case KiloMeter:
		return &Dimensions{
			Length: d.Length * 1.60934,
			Width:  d.Width * 1.60934,
			Height: d.Height * 1.60934,
			Unit:   Mile,
		}
	default:
		return nil
	}
}

// ToKiloMeter converts the dimensions to kilometers
func (d *Dimensions) ToKiloMeter() *Dimensions {
	switch d.Unit {
	case Millimeter:
		return &Dimensions{
			Length: d.Length / 1000000,
			Width:  d.Width / 1000000,
			Height: d.Height / 1000000,
			Unit:   KiloMeter,
		}
	case Centimeter:
		return &Dimensions{
			Length: d.Length / 100000,
			Width:  d.Width / 100000,
			Height: d.Height / 100000,
			Unit:   KiloMeter,
		}
	case Meter:
		return &Dimensions{
			Length: d.Length / 1000,
			Width:  d.Width / 1000,
			Height: d.Height / 1000,
			Unit:   KiloMeter,
		}
	case Inch:
		return &Dimensions{
			Length: d.Length / 39370.1,
			Width:  d.Width / 39370.1,
			Height: d.Height / 39370.1,
			Unit:   KiloMeter,
		}
	case Foot:
		return &Dimensions{
			Length: d.Length / 3280.84,
			Width:  d.Width / 3280.84,
			Height: d.Height / 3280.84,
			Unit:   KiloMeter,
		}
	case Yard:
		return &Dimensions{
			Length: d.Length / 1093.61,
			Width:  d.Width / 1093.61,
			Height: d.Height / 1093.61,
			Unit:   KiloMeter,
		}
	case Mile:
		return &Dimensions{
			Length: d.Length / 1.60934,
			Width:  d.Width / 1.60934,
			Height: d.Height / 1.60934,
			Unit:   KiloMeter,
		}
	case KiloMeter:
		return d
	default:
		return nil
	}
}
