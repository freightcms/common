package graph

import (
	"github.com/freightcms/common/models"
	"github.com/graphql-go/graphql"
)

// EquipmentType is a GraphQL type for the equipment used to transport a shipment.
var EquipmentType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "Equipment",
	Description: "The type of equipment used to transport a shipment.",
	Values: graphql.EnumValueConfigMap{
		"DRY_VAN": &graphql.EnumValueConfig{
			Value: models.DryVan,
		},
		"REEFER": &graphql.EnumValueConfig{
			Value: models.Reefer,
		},
		"FLATBED": &graphql.EnumValueConfig{
			Value: models.Flatbed,
		},
		"STEP_DECK": &graphql.EnumValueConfig{
			Value: models.StepDeck,
		},
		"DOUBLE_DROP": &graphql.EnumValueConfig{
			Value: models.DoubleDrop,
		},
		"LOWBOY": &graphql.EnumValueConfig{
			Value: models.Lowboy,
		},
		"POWER_ONLY": &graphql.EnumValueConfig{
			Value: models.PowerOnly,
		},
		"CONTAINER": &graphql.EnumValueConfig{
			Value: models.Container,
		},
		"CHASSIS": &graphql.EnumValueConfig{
			Value: models.Chassis,
		},
		"TANKER": &graphql.EnumValueConfig{
			Value: models.Tanker,
		},
		"HOPPER": &graphql.EnumValueConfig{
			Value: models.Hopper,
		},
		"PNEUMATIC": &graphql.EnumValueConfig{
			Value: models.Pneumatic,
		},
		"DUMP": &graphql.EnumValueConfig{
			Value: models.Dump,
		},
		"LIVESTOCK": &graphql.EnumValueConfig{
			Value: models.Livestock,
		},
		"AUTO_CARRIER": &graphql.EnumValueConfig{
			Value: models.AutoCarrier,
		},
		"OTHER": &graphql.EnumValueConfig{
			Value: models.Other,
		},
	},
})
