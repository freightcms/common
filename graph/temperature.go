package graph

import (
	"github.com/graphql-go/graphql"
)

var TemperatureUnitGraphQLType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "TemperatureUnit",
	Description: "The unit of measurement for a temperature.",
	Values: graphql.EnumValueConfigMap{
		"CELSIUS": &graphql.EnumValueConfig{
			Value: "CELSIUS",
		},
		"FAHRENHEIT": &graphql.EnumValueConfig{
			Value: "FAHRENHEIT",
		},
	},
})

var TemperatureGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Temperature",
	Description: "The temperature of a product.",
	Fields: graphql.Fields{
		"amount": &graphql.Field{
			Name: "amount",
			Type: graphql.Float,
		},
		"unit": &graphql.Field{
			Name: "unit",
			Type: TemperatureUnitGraphQLType,
		},
	},
})
