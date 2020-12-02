package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongo, err := mongo.NewClient(options.Client().ApplyURI("urltomongo"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongo.Disconnect(ctx)

	api, err := InitializeApi(mongo.Database("default"))
	api.SetupRoutes(gin.Default())
}
