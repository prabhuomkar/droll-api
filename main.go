package main

import (
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/prabhuomkar/droll-api/gql"
	"github.com/prabhuomkar/droll-api/utils"
)

func main() {
	typeDefs, err := gql.GetSchema()

	h := handler.New(&handler.Config{
		Schema:   &typeDefs,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
		log.Printf("Cannot start %s", utils.APIName)
	} else {
		log.Printf("Started %s", utils.APIName)
		log.Printf("Listening on: %s", os.Getenv("PORT"))
	}
}
