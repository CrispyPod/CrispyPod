package main

import (
	"time"

	"crispypod.com/crispypod/controllers"
	"crispypod.com/crispypod/db"
	"crispypod.com/crispypod/graph"
	"crispypod.com/crispypod/helpers"
	"crispypod.com/crispypod/rssfeed"
	"crispypod.com/crispypod/schedule"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

func main() {
	helpers.CheckEnvVariables()
	db.ConnectDatabase()

	r := gin.Default()
	r.Use(helpers.JWTMiddleWare())

	if gin.Mode() == "debug" {
		pH := playground.Handler("GraphQL", "/graphql")
		r.GET("/graphql", func(ctx *gin.Context) {
			pH.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	gH := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.POST("/graphql", func(ctx *gin.Context) {
		gH.ServeHTTP(ctx.Writer, ctx.Request)
	})

	r.StaticFile("/rss", "./Feed/rss.xml")

	r.POST("/api/audioFile", controllers.AudioFileUpload)
	r.GET("/api/audioFile/:fileName", controllers.GetAudioFile)

	rssfeed.GenerateRSSFeed()

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("0:00").Do(schedule.ClearAudioFile)
	s.StartAsync()

	r.Run()
}
