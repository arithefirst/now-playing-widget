package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoServer struct {
	Host       string
	Port       uint16
	Collection string
	DB         string
}

type user struct {
	ID    string `bson:"_id"`
	UID   string `bson:"uid"`
	TC    string `bson:"tc"`
	STC   string `bson:"stc"`
	BG    string `bson:"bg"`
	Empty bool
}

func get(uid string) (*user, error) {
	serverData := mongoServer{
		Host:       "",
		Port:       27017,
		DB:         "config",
		Collection: "users",
	}

	ctx := context.TODO()

	// Set the mongoDB server location
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + serverData.Host + ":" + fmt.Sprint(serverData.Port)))
	if err != nil {
		panic(err)
	}

	// Connect to the server
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// Defer the disconnect to the end of the function
	defer client.Disconnect(ctx)

	// Set the db and collection as well as the search query
	coll := client.Database(serverData.DB).Collection(serverData.Collection)
	filter := bson.D{{Key: "uid", Value: uid}}

	// Find one document matching the input uid
	var output user
	err = coll.FindOne(ctx, filter).Decode(&output)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found with the given name")
			return &user{ID: "", UID: "", TC: "", STC: "", BG: "", Empty: true}, nil
		} else {
			log.Fatal(err)
		}
	}

	output.Empty = false
	return &output, nil
}
