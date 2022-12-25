package model

import (
	"github.com/google/uuid"
)

type Product struct {
	id          uuid.UUID
	name        string
	discription string
}

func CreateProduct(name string, discriprion string) *Product {
	return &Product{
		id:          uuid.New(),
		name:        name,
		discription: discriprion,
	}
}

func (product Product) GetUUID() uuid.UUID {
	return product.id
}
