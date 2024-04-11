package graph

import (
	"github.com/graphql-go/graphql"
)

var ContactGraphQLType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Contact",
	Description: "A contact that can be used to reach a company, organization, or individual.",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"phone": &graphql.Field{
			Type: graphql.String,
		},
		"fax": &graphql.Field{
			Type: graphql.String,
		},
	},
})
