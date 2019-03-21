package main

import (
	"context"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jkieltyka/graphql-practice/mongo"
	"github.com/jkieltyka/graphql-practice/service"
	"google.golang.org/grpc/resolver"
)

type query struct{}

func (_ *query) Name() string { return "Hello, world!" }

func (_ *query) Age() int32 { return 10 }

func main() {
	ctx := context.Background()
	db, _ := mongo.ConfigDB(ctx)
	//todo add error handling here
	roleService := service.NewNodeService(db)
	ctx.WithValue("roleService", roleService)

	schema := graphql.MustParseSchema(schema.GetRoot(), &resolver.Resolver{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
