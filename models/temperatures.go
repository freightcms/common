package models

// Temperature is a struct that represents a temperature value and its unit
type TemperatureUnit string

const (
	Celsius    TemperatureUnit = "C"
	Fahrenheit TemperatureUnit = "F"
)

type Temperature struct {
	Value float64         `json:"value" bson:"value"` // The temperature value
	Unit  TemperatureUnit `json:"unit" bson:"unit"`   // The unit of the temperature
}

// Creates a copy of the current temperature in Celsius
func (t *Temperature) ToCelsius() *Temperature {
	if t.Unit == Fahrenheit {
		return &Temperature{
			Value: (t.Value - 32) * 5 / 9,
			Unit:  Celsius,
		}
	}
	return &Temperature{
		Value: t.Value,
		Unit:  t.Unit,
	}
}

// Creates a copy of the current temperature in Fahrenheit
func (t *Temperature) ToFahrenheit() *Temperature {
	if t.Unit == Celsius {
		return &Temperature{
			Value: t.Value*9/5 + 32,
			Unit:  Fahrenheit,
		}
	}
	return &Temperature{
		Value: t.Value,
		Unit:  t.Unit,
	}
}
