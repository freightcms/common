package graph

import (
	"github.com/freightcms/common/models"
	"github.com/graphql-go/graphql"
)

// DimensionUnitGraphQLType is a GraphQL type for the unit of measurement for a dimension.
var DimensionUnitGraphQLType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "DimensionUnit",
	Description: "The unit of measurement for a dimension.",
	Values: graphql.EnumValueConfigMap{
		"MILLIMETERS": &graphql.EnumValueConfig{
			Value: models.Millimeter,
		},
		"CENTIMETERS": &graphql.EnumValueConfig{
			Value: models.Centimeter,
		},
		"METERS": &graphql.EnumValueConfig{
			Value: models.Meter,
		},
		"INCHES": &graphql.EnumValueConfig{
			Value: models.Inch,
		},
		"FOOT": &graphql.EnumValueConfig{
			Value: models.Foot,
		},
		"YARDS": &graphql.EnumValueConfig{
			Value: models.Yard,
		},
		"MILES": &graphql.EnumValueConfig{
			Value: models.Mile,
		},
		"KILOMETERS": &graphql.EnumValueConfig{
			Value: models.KiloMeter,
		},
	},
})

// DimensionsGraphQLType is a GraphQL type for the dimensions of a product.
var DimensionsGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Dimensions",
	Description: "The dimensions of a product.",
	Fields: graphql.Fields{
		"length": &graphql.Field{
			Name: "length",
			Type: graphql.Float,
		},
		"width": &graphql.Field{
			Name: "width",
			Type: graphql.Float,
		},
		"height": &graphql.Field{
			Name: "height",
			Type: graphql.Float,
		},
		"unit": &graphql.Field{
			Name: "unit",
			Type: DimensionUnitGraphQLType,
		},
	},
})
