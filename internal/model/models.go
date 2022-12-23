package model

import (
	"github.com/google/uuid"
)

type Product struct {
	id          uuid.UUID
	name        string
	discription string
}

type User struct {
	id   uuid.UUID
	name string
}

type Order struct {
	id          uuid.UUID
	location    string
	productList map[string]int
}

type Cart struct {
	id                   uuid.UUID
	user                 User
	productNameCountPair map[string]int
}

func (product Product) GetProductUUID() uuid.UUID {
	return product.id
}

func (user User) GetUUID() uuid.UUID {
	return user.id
}

func (order Order) GetUUID() uuid.UUID {
	return order.id
}

func (cart Cart) GetUUID() uuid.UUID {
	return cart.id
}

func (product *Product) Update() {

}

func (user *User) Update() {

}

func (order *Order) Update() {

}

func (cart *Cart) Update() {

}
