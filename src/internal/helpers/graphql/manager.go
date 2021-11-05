package graphql

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ppcamp/go-graphql-with-auth/internal/helpers/controller"
)

type Schema struct {
	queries   graphql.Fields
	mutations graphql.Fields

	authQueries   graphql.Fields
	authMutations graphql.Fields
}

func NewSchemaManager() *Schema {
	return &Schema{
		queries:       graphql.Fields{},
		mutations:     graphql.Fields{},
		authQueries:   graphql.Fields{},
		authMutations: graphql.Fields{},
	}
}

func (s *Schema) registerAuthQueries() {
	if len(s.authQueries) > 0 {
		var me = graphql.NewObject(
			graphql.ObjectConfig{
				Name:        "MeQuery",
				Description: "Type to encapsulate all authenticated queries",
				Fields:      s.authQueries,
			},
		)

		s.queries["me"] = &graphql.Field{
			Type:        me,
			Description: "Run some query with jwt auth",
			Resolve:     controller.AuthorizedOnly,
		}
	}
}

func (s *Schema) registerAuthMutations() {
	if len(s.authMutations) > 0 {
		var me = graphql.NewObject(
			graphql.ObjectConfig{
				Name:        "MeMutation",
				Description: "Type to encapsulate all authenticated queries",
				Fields:      s.authMutations,
			},
		)

		s.mutations["me"] = &graphql.Field{
			Type:        me,
			Description: "Run some mutation with jwt auth",
			Resolve:     controller.AuthorizedOnly,
		}
	}
}

func (s *Schema) RegisterQuery(fieldName string, fieldValue *graphql.Field) {
	s.queries[fieldName] = fieldValue
}

func (s *Schema) RegisterMutation(fieldName string, fieldValue *graphql.Field) {
	s.mutations[fieldName] = fieldValue
}

func (s *Schema) RegisterAuthenticatedQuery(fieldName string, fieldValue *graphql.Field) {
	s.authQueries[fieldName] = fieldValue
}

func (s *Schema) RegisterAuthenticatedMutation(fieldName string, fieldValue *graphql.Field) {
	s.authMutations[fieldName] = fieldValue
}

func (s *Schema) getSchemas() graphql.Schema {
	s.registerAuthQueries()
	s.registerAuthMutations()

	// Schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Query",
			Description: "All elements that can be fetched",
			Fields:      s.queries,
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Mutation",
			Description: "All functions that make some change in API",
			Fields:      s.mutations,
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
	schema := s.getSchemas()

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
		// Playground: true,
	})

	return func(c *gin.Context) {
		h.ContextHandler(c, c.Writer, c.Request)
	}
}
