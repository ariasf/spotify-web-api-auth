package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"log"
	"main/spotify"
	"net/http"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "it be good 🧐💖")
	})

	r.GET("spotify/login", func(context *gin.Context) {
		spotify.Login(context)
	})
	r.POST("spotify/swap", func(context *gin.Context) {
		spotify.Swap(context)
	})
	r.POST("spotify/refreshtoken", func(context *gin.Context) {
		spotify.Refresh(context)
	})

	ginLambda = ginadapter.New(r)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)

}
