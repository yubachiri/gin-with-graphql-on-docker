package queries

import (
	"github.com/graphql-go/graphql"
)

// Hello basic query
var Hello graphql.Field = graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "var", nil
	},
}

// HelloWithArg query with argument "id"
var HelloWithArg graphql.Field = graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(string)
		if isOK {
			return idQuery, nil
		}

		return "", nil
	},
}
