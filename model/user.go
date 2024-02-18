package usermodel

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func GetUser() ([]map[string]interface{}, error) {
    var users []map[string]interface{}
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
        // Convert User struct to map[string]interface{}
        userMap := map[string]interface{}{
            "first_name": user.FirstName,
            "last_name":  user.LastName,
            "email":      user.Email,
            "gender":     user.Gender,
        }
        users = append(users, userMap)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}