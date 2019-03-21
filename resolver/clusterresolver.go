package resolver

import "github.com/jkieltyka/graphql-practice/model"

type clusterResolver struct {
	cluster *model.Cluster
}

func (r *clusterResolver) Name() *string {
	return &r.cluster.Name
}

func (r *clusterResolver) Nodes() *[]*nodeResolver {
	n := make([]*nodeResolver, len(r.cluster.Nodes))
	for i := range n {
		n[i] = &nodeResolver{
			node: r.cluster.Nodes[i],
		}
	}
	return &n
}
