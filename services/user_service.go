package services

import (
	"github.com/google/uuid"
	"go-crud/models"
)
// Create a New User
func NewUser(name, email string, age int) *models.User {
	return &models.User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
		Age:   age,
	}
}
