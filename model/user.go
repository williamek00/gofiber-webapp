package usermodel

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	database "learngo/db"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Gender    string             `bson:"gender"`
}

var db, _ = database.Db()
var userCollection = db.Collection("admin")

func GetUsers() ([]User, error) {
	var users []User
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

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(c *fiber.Ctx) User {
	var user User

	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error converting ID to ObjectID:", err)
		return User{}
	}

	ctx := context.TODO()
	filter := bson.D{{Key: "_id", Value: objID}}

	err = userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return User{}
	}

	return user
}

func CreateUser(c *fiber.Ctx) string {
	var user = User{}

	if err := c.BodyParser(&user); err != nil {
		panic(err)
	}

	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	return "Successfuly add new user"
}

func DeleteUser(c *fiber.Ctx) string {
	id := c.Params("id")

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = userCollection.DeleteOne(ctx, filter)

	if err != nil {
		panic(err)
	}
	msg := fmt.Sprintf("User with id %d", objID)
	return msg
}
