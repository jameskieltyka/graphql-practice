package service

import (
	"github.com/jkieltyka/graphql-practice/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func NewNodeService(db *mongo.Database, ctx context.Context) *NodeService {
	return &NodeService{db, ctx}
}

type NodeService struct {
	db  *mongo.Database
	ctx context.Context
}

func (n *NodeService) AddNode(node *model.Node) (*model.Node, error) {

	res, err := n.db.Collection("nodes").InsertOne(n.ctx, bson.D{{"name", &node.Name}})
	if err != nil {
		return nil, err
	}

	node.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return node, err
}

func (n *NodeService) GetNode(name string) (res *model.Node, err error) {
	nameDoc := bson.D{{"name", name}}
	err = n.db.Collection("nodes").FindOne(n.ctx, nameDoc).Decode(res)
	return res, err
}
