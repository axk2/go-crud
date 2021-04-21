package main

import (
	"time"
	m "github.com/axk2/crud/models"
	userService "github.com/axk2/crud/services/user.service"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	user := m.User{
		ID:			primitive.NewObjectID(),
		Name:  		"Mario",
		Email: 		"mario@mario.com",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	userService.Create(user)
}