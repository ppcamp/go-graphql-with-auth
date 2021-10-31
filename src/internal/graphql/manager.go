package graphql

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Schema struct {
	queries   graphql.Fields
	mutations graphql.Fields
}

func NewSchemaManager() *Schema {
	return &Schema{
		queries:   graphql.Fields{},
		mutations: graphql.Fields{},
	}
}

func (s *Schema) RegisterQuery(fieldName string, fieldValue *graphql.Field) {
	s.queries[fieldName] = fieldValue
}

func (s *Schema) RegisterMutation(fieldName string, fieldValue *graphql.Field) {
	s.mutations[fieldName] = fieldValue
}

func (s *Schema) Get() graphql.Schema {
	// Schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: s.queries,
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: s.mutations,
		}),
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

// Handler is a closure that will wrap the schema and return a proper gin
// handler
func (s *Schema) Handler() gin.HandlerFunc {
	schema := s.Get()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
		// Playground: true,
	})

	return gin.WrapH(h)

}
