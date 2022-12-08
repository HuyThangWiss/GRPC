package main

import (
	"B/Grpc_Handler"
	"B/adapters/database"
	"B/core/Services"
	"B/core/ports"
	"B/pb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	mongoCollection := NewMongoDBCollection()
	mongoDB := database.NewMonGoDb(mongoCollection)
	PostRepository := ports.InitPostRepositoryPort(mongoDB)
	PostSerives := Services.NewUserService(PostRepository)
	PostServer := Grpc_Handler.NewGrpcServer(PostSerives)
	s:=grpc.NewServer()
	pb.RegisterPostServicesServer(s,PostServer)
	err = s.Serve(lis)

	if err != nil{
		fmt.Println("err ",err)
	}

}
func NewMongoDBCollection() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://admin:f6XPinsVTx@localhost:27017/?readPreference=primary&directConnection=true&ssl=false"))
	if err != nil {
		log.Fatalf("NewMongoDB err: %v", err)
		return nil
	}
	db := client.Database("Posts")
	collection := db.Collection("Post")
	return collection
}