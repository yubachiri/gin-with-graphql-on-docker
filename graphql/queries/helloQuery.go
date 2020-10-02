package queries

import "github.com/graphql-go/graphql"

func Hello() graphql.Field {
	field := graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	}

	return field
}
