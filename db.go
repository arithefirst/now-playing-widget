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
	UID   string `bson:"uid"`
	TC    string `bson:"tc"`
	STC   string `bson:"stc"`
	BG    string `bson:"bg"`
	RIGHT bool   `bson:"right"`
	Empty bool
}

func get(uid string) (*user, error) {
	serverData := mongoServer{
		Host:       "10.0.0.21",
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
			fmt.Printf("No documents found with the id \"%v\"\n", uid)
			return &user{UID: "", TC: "", STC: "", BG: "", RIGHT: false, Empty: true}, nil
		} else {
			log.Fatal(err)
		}
	}

	output.Empty = false
	return &output, nil
}

func set(user user) (*mongo.UpdateResult, error) {
	serverData := mongoServer{
		Host:       "10.0.0.21",
		Port:       27017,
		DB:         "config",
		Collection: "users",
	}

	ctx := context.Background()
	opts := options.Update().SetUpsert(true)

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

	// Set the db and collection
	db := client.Database(serverData.DB)
	collection := db.Collection(serverData.Collection)

	// Set the filter
	var filter bson.D = bson.D{{Key: "uid", Value: user.UID}}

	// Set the data to be sent to mongoDB
	var update bson.D = bson.D{
		{Key: "bg", Value: user.BG},
		{Key: "tc", Value: user.TC},
		{Key: "stc", Value: user.STC},
		{Key: "uid", Value: user.UID},
		{Key: "right", Value: user.RIGHT},
	}

	result, err := collection.UpdateOne(
		ctx,
		filter,
		bson.M{"$set": update},
		opts,
	)

	return result, err

}
