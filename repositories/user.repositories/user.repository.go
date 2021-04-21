package user_repository

import (
	"github.com/axk2/crud/database"
	m "github.com/axk2/crud/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("usuarios")
var ctx = context.Background()

func Create(user m.User) error {
	var err error
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Users, error) {
	var users m.Users
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user m.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func Update(user m.User, userId string) error {
	return nil
}

func Delete(userId string) error {
	return nil
}