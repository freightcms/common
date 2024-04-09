package graph

import (
	"github.com/graphql-go/graphql"
)

var CurrencyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Currency",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Name: "code",
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
		},
		"symbol": &graphql.Field{
			Name: "symbol",
			Type: graphql.String,
		},
	},
})
