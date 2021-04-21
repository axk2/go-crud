package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Datos de Usuario
type User struct {
	ID			primitive.ObjectID	`bson:"_id"`
	Name 		string				`bson:"name"`
	Email		string				`bson:"email"`
	CreatedAt	time.Time			`bson:"created_at"`
	UpdatedAt	time.Time			`bson:"updated_at,omitempty"`
}

// Lista de Usuarios
type Users []*User