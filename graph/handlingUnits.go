package graph

import "github.com/graphql-go/graphql"

// HandlingUnitType is a GraphQL type for a handling unit.
var HandlingUnitType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "HandlingUnit",
	Description: "A handling unit.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"weight": &graphql.Field{
			Type: WeightGraphQLType,
		},
		"dimensions": &graphql.Field{
			Type: DimensionsGraphQLType,
		},
		"volume": &graphql.Field{
			Type: VolumeGraphQLType,
		},
		"nfmc": &graphql.Field{
			Type: graphql.String,
		},
		"countryOfOrigin": &graphql.Field{
			Type: graphql.String,
		},
		"declaredValue": &graphql.Field{
			Type: CurrencyGraphQLType,
		},
		"isHazardous": &graphql.Field{
			Type: graphql.Boolean,
		},
		"isStackable": &graphql.Field{
			Type: graphql.Boolean,
		},
		"metadata": &graphql.Field{
			Type: MetadataGraphQLType,
		},
		"minTemperature": &graphql.Field{
			Type: TemperatureGraphQLType,
		},
		"maxTemperature": &graphql.Field{
			Type: TemperatureGraphQLType,
		},
		"references": &graphql.Field{
			Type: graphql.NewList(ReferenceGraphQLType),
		},
		"commodities": &graphql.Field{
			Type: graphql.NewList(CommodityGraphQLType),
		},
	},
})
