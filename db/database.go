package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	// "log"
    // "time"
)

// func Db() *mongo.Database {
//     var db *mongo.Database
//     var ctx = context.TODO()

//     client, connErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://williamenosk123:BaEDDLx1qry6VpKk@cluster0.9agejbd.mongodb.net"))

//     if connErr != nil {
//         log.Fatal(connErr)
//     }
//     //TODO : magic string
//     db = client.Database("learngo-db")

//     return db
// }

func Db() (*mongo.Database, error) {
    var ctx = context.TODO()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://williamenosk123:BaEDDLx1qry6VpKk@cluster0.9agejbd.mongodb.net"))
    if err != nil {
        return nil, err
    }

    // Check if the connection was successful
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    return client.Database("learngo-db"), nil
}