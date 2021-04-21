package user_service_test

import (
	"testing"
	userService "github.com/axk2/crud/services/user.service"
	m "github.com/axk2/crud/models"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func TestCreate(t *testing.T) {
	user := m.User{
		ID:			primitive.NewObjectID(),
		Name:  		"Jesus",
		Email: 		"jesus@jesus.com",
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}
	err := userService.Create(user)
	if err != nil {
		t.Error("Prueba de persistencia de datos fallida")
	} else {
		t.Log("Prueba finalizada exitosamente")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()
	if err != nil {
		t.Error("Error en lectura")
	}
	if len(users) == 0 {
		t.Error("No hay datos")
	} else {
		t.Log("todo bien")
	}
}

func TestUpdate(t *testing.T) {
	
}

func TestDelete(t *testing.T) {
	
}