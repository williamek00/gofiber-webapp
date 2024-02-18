package usermodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"context"
	// "log"
	"learngo/db"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Gender    string             `bson:"gender"`
}
 
var db,_ = database.Db()
var userCollection = db.Collection("admin")

func GetUser() ([]string, error) {
    var users []string
    var ctx = context.TODO()

    filter := bson.D{{}}

    cursor, findErr := userCollection.Find(ctx, filter)
    if findErr != nil {
        return nil, findErr
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        data, marshalErr := bson.MarshalExtJSON(user, false, false)
        if marshalErr != nil {
            return nil, marshalErr
        }
        users = append(users, string(data))
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}