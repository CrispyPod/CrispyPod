package main

import (
	"crispypod.com/crispypod/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pH := playground.Handler("GraphQL", "/graphql")
	r.GET("/graphql", func(ctx *gin.Context) {
		pH.ServeHTTP(ctx.Writer, ctx.Request)
	})

	gH := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.POST("/graphql", func(ctx *gin.Context) {
		gH.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.Run()
}
