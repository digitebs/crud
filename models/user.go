package models

import (
	"context"
	"crud/config"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
"id": "xxx", // user ID
"name": "test", // user name
"dob": "", // date of birth
"address": "", // user address
"description": "", // user description
"createdAt": "" // user created date }
*/

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string `json:"name"`
	DOB         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

func init() {
	config.Connect()
	collection = config.GetCollection()
}

var collection *mongo.Collection

/*
	create user
*/
func (user User) CreateUser() (User, error) {
	user.ID = primitive.NewObjectID()
	updateResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	fmt.Printf("Created User %v", updateResult.InsertedID)
	return user, err
}

/*
	update user

 */
func UpdateUser(user User) error {

	filter := bson.D{{"_id", user.ID}}

	updateResult, err := collection.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if updateResult.MatchedCount == 0 {
		fmt.Println("no document match")
		return  errors.New("no document match")
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

/*
	get user by id
 */
func GetUserById(id string) (User, error) {
	var result User
	objectid, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		fmt.Println(err2)
		return result, err2
	}
	filter := bson.D{{"_id",objectid}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, nil
}

/*
	delete by user id
 */
func DeleteUserById(id string) error {
	objectid, err2 := primitive.ObjectIDFromHex(id)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}
	filter := bson.D{{"_id", objectid}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted %v documents in collection\n", deleteResult.DeletedCount)
	return nil
}
