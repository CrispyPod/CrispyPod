package main

import (
	"crispypod.com/crispypod/controllers"
	"crispypod.com/crispypod/graph"
	"crispypod.com/crispypod/helpers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	helpers.CheckEnvVariables()

	r := gin.Default()

	pH := playground.Handler("GraphQL", "/graphql")
	r.GET("/graphql", func(ctx *gin.Context) {
		pH.ServeHTTP(ctx.Writer, ctx.Request)
	})

	gH := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.POST("/graphql", func(ctx *gin.Context) {
		gH.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.POST("/login",controllers.Login

	r.Run()
}
