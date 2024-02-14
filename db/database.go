package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	"fmt"
	"log"
)

func Init() string {
    // Initialize variables
    var collection *mongo.Collection
    var ctx = context.TODO()

    // Define struct to hold query results
    var result struct {
        ID        primitive.ObjectID `bson:"_id"`
        FirstName string             `bson:"first_name"`
        LastName  string             `bson:"last_name"`
        Email     string             `bson:"email"`
        Gender    string             `bson:"gender"`
    }

    // Connect to MongoDB
    client, connErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://williamenosk123:BaEDDLx1qry6VpKk@cluster0.9agejbd.mongodb.net"))

    if connErr != nil {
        log.Fatal(connErr)
    }

    // Specify the database and collection
    collection = client.Database("learngo-db").Collection("admin")

    // Define an empty filter to retrieve all documents
    filter := bson.D{{}}

    // Find documents in the collection
    cursor, findErr := collection.Find(ctx, filter)
    if findErr != nil {
        log.Fatal(findErr)
    }
    defer cursor.Close(ctx) // Make sure to close the cursor when done

    // Iterate through the cursor and decode each document into the 'result' variable
    for cursor.Next(ctx) {
        if err := cursor.Decode(&result); err != nil {
            log.Fatal(err)
        }
    }

    // Check if there were any errors during the cursor iteration
    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }

    // Marshal the result into JSON format
    data, marshalErr := bson.MarshalExtJSON(result, false, false)
    if marshalErr != nil {
        log.Fatal(marshalErr)
    }

    // Return the JSON data as a string
    return string(data)
}
