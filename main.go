package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/prabhuomkar/droll-api/gql"
)

func main() {
	fields := graphql.Fields{
		"version": &graphql.Field{
			Type:    gql.VersionType,
			Resolve: gql.VersionQueryResolver,
		},
		"xkcd": &graphql.Field{
			Type:    graphql.NewList(gql.XKCDType),
			Args:    gql.Args,
			Resolve: gql.XKCDQueryResolver,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		log.Printf("Cannot start %s", gql.Name)
	} else {
		log.Printf("Started %s", gql.Name)
		log.Printf("Listening on: %s", os.Getenv("PORT"))
	}
}
