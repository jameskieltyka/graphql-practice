package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jkieltyka/graphql-practice/model"
)

type nodeResolver struct {
	node *model.Node
}

func (r *nodeResolver) Name() *string {
	return &r.node.Name
}

func (r *nodeResolver) ID() graphql.ID {
	return graphql.ID(r.node.ID)
}
