package queries

import (
	"fmt"

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
