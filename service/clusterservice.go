package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type ClusterService struct {
	db          *mongo.Database
	nodeService *NodeService
	ctx         context.Context
}
