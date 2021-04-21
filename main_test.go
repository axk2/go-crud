package main

import (
	"testing"
	m "github.com/axk2/crud/models"
	userService "github.com/axk2/crud/services/user.service"
	"time"
)

func MainTest(t *testing.T) {
	user := m.User{
		Name:  		"Mario",
		Email: 		"mario@mario.com",
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