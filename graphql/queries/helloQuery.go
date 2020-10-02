package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// Hello 動く
func Hello() graphql.Field {
	field := graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	}

	return field
}

// HelloArg 動かない
func HelloArg() graphql.Field {
	field := graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			idQuery, isOK := p.Args["id"].(string)
			fmt.Printf(idQuery)
			if isOK {
				return idQuery, nil
			}

			return "", nil
		},
	}

	return field
}
