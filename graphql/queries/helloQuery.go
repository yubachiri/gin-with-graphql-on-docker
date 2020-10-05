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

// HelloVar HelloField
var HelloVar graphql.Field = graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "var", nil
	},
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
			fmt.Printf("%v\n", p)
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

// HelloVarArg HelloField
var HelloVarArg graphql.Field = graphql.Field{
	Type:        graphql.String,
	Description: "Print Arg",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		fmt.Printf("%v\n", p)
		idQuery, isOK := p.Args["id"].(string)
		fmt.Printf(idQuery)
		if isOK {
			return idQuery, nil
		}

		return "", nil
	},
}
