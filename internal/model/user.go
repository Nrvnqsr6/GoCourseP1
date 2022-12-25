package model

import (
	"github.com/google/uuid"
)

type User struct {
	id   uuid.UUID
	name string
}

func CreateUser(name string) *User {
	return &User{
		id:   uuid.New(),
		name: name,
	}
}

func (user *User) Update(newName string) {
	user.name = newName
}

func (user User) GetUUID() uuid.UUID {
	return user.id
}
