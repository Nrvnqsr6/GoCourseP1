package model

import (
	"github.com/google/uuid"
)

type Order struct {
	id       uuid.UUID
	location string
	user     *User
}

func CreateOrder(location string, user *User) *Order {
	return &Order{
		id:       uuid.New(),
		location: location,
		user:     user,
	}
}

func (order Order) GetUUID() uuid.UUID {
	return order.id
}
