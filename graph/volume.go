package graph

import (
	"github.com/freightcms/common/models"
	"github.com/graphql-go/graphql"
)

// VolumeUnitGraphQLType is a GraphQL type for the unit of measurement for a volume.
var VolumeUnitGraphQLType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "VolumeUnit",
	Description: "The unit of measurement for a volume.",
	Values: graphql.EnumValueConfigMap{
		"CUBIC_METRES": &graphql.EnumValueConfig{
			Value: models.CubicMetre,
		},
		"CUBIC_INCHES": &graphql.EnumValueConfig{
			Value: models.CubicInch,
		},
		"CUBIC_FEET": &graphql.EnumValueConfig{
			Value: models.CubicFoot,
		},
		"CUBIC_YARDS": &graphql.EnumValueConfig{
			Value: models.CubicYard,
		},
		"CUBIC_DECIMETRES": &graphql.EnumValueConfig{
			Value: models.CubicDecimetre,
		},
	},
})

// VolumeGraphQLType is a GraphQL type for the volume of a product.
var VolumeGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Volume",
	Description: "The volume of a product.",
	Fields: graphql.Fields{
		"amount": &graphql.Field{
			Name: "amount",
			Type: graphql.Float,
		},
		"unit": &graphql.Field{
			Name: "unit",
			Type: VolumeUnitGraphQLType,
		},
	},
})
