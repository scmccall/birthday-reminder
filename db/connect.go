package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

	"time"
	"log"
	"context"
	"fmt"
)

func ConnectToDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(getURI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("people")
	peopleCollection := database.Collection("info")

	cursor, err := peopleCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var people []bson.M
	if err = cursor.All(ctx, &people); err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)

}