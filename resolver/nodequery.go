package resolver

import "context"

func (r *resolver) GetNode(ctx context.Context, args struct {
	Name string
})
