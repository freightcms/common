package graph

import (
	"github.com/freightcms/common/models"
	"github.com/graphql-go/graphql"
)

// WeightUnitGraphQLType is a GraphQL type for the unit of measurement for a weight.
var WeightUnitGraphQLType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "WeightUnit",
	Description: "The unit of measurement for a weight.",
	Values: graphql.EnumValueConfigMap{
		"GRAMS": &graphql.EnumValueConfig{
			Value: models.Gram,
		},
		"KILOGRAMS": &graphql.EnumValueConfig{
			Value: models.Kilogram,
		},
		"OUNCES": &graphql.EnumValueConfig{
			Value: models.Ounce,
		},
		"POUNDS": &graphql.EnumValueConfig{
			Value: models.Pound,
		},
	},
})

// WeightGraphQLType is a GraphQL type for the weight of a product.
var WeightGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Weight",
	Description: "The weight of a product.",
	Fields: graphql.Fields{
		"value": &graphql.Field{
			Name: "value",
			Type: graphql.Float,
		},
		"unit": &graphql.Field{
			Name: "unit",
			Type: WeightUnitGraphQLType,
		},
	},
})
