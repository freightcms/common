package graph

import (
	"github.com/graphql-go/graphql"
)

// ReferenceType is a GraphQL type for a reference to a document or other resource.
var ReferenceGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Reference",
	Description: "A reference to a document or other resource.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"reference": &graphql.Field{
			Type: graphql.String,
		},
		"source": &graphql.Field{
			Type: graphql.String,
		},
	},
})
