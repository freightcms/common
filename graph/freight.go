package graph

import (
	"github.com/freightcms/common/models"
	"github.com/graphql-go/graphql"
)

// FreightType is a GraphQL type for the type of freight being transported.
var FreightType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Freight",
	Description: "The type of freight being transported.",
	Values: graphql.EnumValueConfigMap{
		"CLASS_50": &graphql.EnumValueConfig{
			Value: models.Class50,
		},
		"CLASS_55": &graphql.EnumValueConfig{
			Value: models.Class55,
		},
		"CLASS_60": &graphql.EnumValueConfig{
			Value: models.Class60,
		},
		"CLASS_65": &graphql.EnumValueConfig{
			Value: models.Class65,
		},
		"CLASS_70": &graphql.EnumValueConfig{
			Value: models.Class70,
		},
		"CLASS_77_55": &graphql.EnumValueConfig{
			Value: models.Class77_55,
		},
		"CLASS_85": &graphql.EnumValueConfig{
			Value: models.Class85,
		},
		"CLASS_92_5": &graphql.EnumValueConfig{
			Value: models.Class92_5,
		},
		"CLASS_100": &graphql.EnumValueConfig{
			Value: models.Class100,
		},
		"CLASS_110": &graphql.EnumValueConfig{
			Value: models.Class110,
		},
		"CLASS_125": &graphql.EnumValueConfig{
			Value: models.Class125,
		},
		"CLASS_150": &graphql.EnumValueConfig{
			Value: models.Class150,
		},
		"CLASS_175": &graphql.EnumValueConfig{
			Value: models.Class175,
		},
		"CLASS_200": &graphql.EnumValueConfig{
			Value: models.Class200,
		},
		"CLASS_250": &graphql.EnumValueConfig{
			Value: models.Class250,
		},
		"CLASS_300": &graphql.EnumValueConfig{
			Value: models.Class300,
		},
		"CLASS_400": &graphql.EnumValueConfig{
			Value: models.Class400,
		},
		"CLASS_500": &graphql.EnumValueConfig{
			Value: models.Class500,
		},
	},
})
