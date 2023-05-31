package main

import (
	"context"
	"example/BatteryTracking/controller"
	"example/BatteryTracking/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server 		*gin.Engine
	us			service.SchoolService
	uc			controller.SchoolController
	ctx			context.Context
	userc		*mongo.Collection
	mongoclient	*mongo.Client
	err			error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connectting with mongo", err)
	}
	err =mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("Goschool").Collection("school")
	us = service.NewSchoolService(userc, ctx)
	uc = controller.New(us)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterUserRouts(basepath)

	log.Fatal(server.Run(":9091"))
}