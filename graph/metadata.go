package graph

// Metadata

import (
	"github.com/graphql-go/graphql"
)

// MetadataGraphQLType is a GraphQL type for metadata.
var MetadataGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Metadata",
	Description: "Metadata for a document or other resource.",
	Fields: graphql.Fields{
		"orderNumber": &graphql.Field{
			Type: graphql.String,
		},
		"invoiceNumbeR": &graphql.Field{
			Type: graphql.String,
		},
		"purchaseOrderNumber": &graphql.Field{
			Type: graphql.String,
		},
		"customerOrderNumber": &graphql.Field{
			Type: graphql.String,
		},
		"billOfLadingNumber": &graphql.Field{
			Type: graphql.String,
		},
		"lineItemNumber": &graphql.Field{
			Type: graphql.String,
		},
		"referenceNumber": &graphql.Field{
			Type: graphql.String,
		},
		"containerNumber": &graphql.Field{
			Type: graphql.String,
		},
		"shipmentNumber": &graphql.Field{
			Type: graphql.String,
		},
		"sealNumber": &graphql.Field{
			Type: graphql.String,
		},
		"bookingNumber": &graphql.Field{
			Type: graphql.String,
		},
		"masterBillOfLadingNumber": &graphql.Field{
			Type: graphql.String,
		},
		"houseBillOfLadingNumber": &graphql.Field{
			Type: graphql.String,
		},
		"proNumber": &graphql.Field{
			Type: graphql.String,
		},
		"trackingNumber": &graphql.Field{
			Type: graphql.String,
		},
		"packageNumber": &graphql.Field{
			Type: graphql.String,
		},
	},
})
