package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/prabhuomkar/droll-api/resolvers"
	"github.com/prabhuomkar/droll-api/schema"
)

func main() {
	fields := graphql.Fields{
		"version": &graphql.Field{
			Type:    schema.VersionType,
			Resolve: resolvers.VersionQueryResolver,
		},
		"xkcd": &graphql.Field{
			Type:    graphql.NewList(schema.XKCDType),
			Args:    schema.Args,
			Resolve: resolvers.XKCDQueryResolver,
		},
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery", Fields: fields,
		}),
	}

	typeDefs, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &typeDefs,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
		log.Printf("Cannot start %s", schema.APIName)
	} else {
		log.Printf("Started %s", schema.APIName)
		log.Printf("Listening on: %s", os.Getenv("PORT"))
	}
}
