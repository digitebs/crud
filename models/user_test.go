package models

import (
	"github.com/mongo-go/testdb"
	"testing"
	"time"
)


func setUp(t *testing.T){
	testDb := testdb.NewTestDB("mongodb://localhost:27017", "test", time.Duration(10) * time.Second)
	testDb.Connect()
	coll, err := testDb.CreateRandomCollection(testdb.NoIndexes)
	if err != nil {
		t.Fatal(err)
	}
	collection = coll
}


var u = User{
}

func TestDeleteUserById(t *testing.T) {
	setUp(t)
	user,err := u.CreateUser()
	if err !=nil {
		t.Error("Fail to create user")
	}
	err2 := DeleteUserById(user.ID.Hex())
	if err2 !=nil {
		t.Error("Fail to delete user")
	}
}

func TestGetUserById(t *testing.T) {
	setUp(t)
	user,err := u.CreateUser()
	if err !=nil {
		t.Error("Fail to create user")
	}
	_,err2 := GetUserById(user.ID.Hex())
	if err2 != nil {
		t.Error("Fail to get user")
	}
}

func TestUpdateUser(t *testing.T) {
	setUp(t)
	user,err := u.CreateUser()
	if(err !=nil){
		t.Error("Fail to create user")
	}
	user.Name ="peter"
	err2 := UpdateUser(user)
	if(err2 !=nil){
		t.Error("Fail to update user")
	}

}

func TestUser_CreateUser(t *testing.T) {
	setUp(t)
	_,err := u.CreateUser()
	if err !=nil {
		t.Error("Fail to create user")
	}
}
