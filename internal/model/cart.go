package model

import (
	"github.com/google/uuid"
)

type Cart struct {
	id                   uuid.UUID
	user                 *User
	productNameCountPair map[string]int
}

func CreateCart(user *User) *Cart {
	return &Cart{
		id:                   uuid.New(),
		user:                 user,
		productNameCountPair: make(map[string]int),
	}

}

// так конечно делать нехорошо, но прописывать отдельно методы обновления каждого поля лень
func (cart *Cart) Update(newCart *Cart) {
	if newCart.user != nil {
		cart.user = newCart.user
	}
	if newCart.productNameCountPair != nil {
		cart.productNameCountPair = newCart.productNameCountPair
	}
}

func (cart *Cart) GetUUID() uuid.UUID {
	return cart.id
}

// package model

// import (
// 	"github.com/google/uuid"
// )

// type Cart struct {
// 	id                   uuid.UUID
// 	User                 *User
// 	ProductNameCountPair map[string]int
// }

// func CreateCart(user *User) *Cart {
// 	return &Cart{
// 		id:                   uuid.New(),
// 		User:                 user,
// 		ProductNameCountPair: make(map[string]int),
// 	}

// }

// func (cart *Cart) Update(newCart *Cart) {
// 	if newCart.User != nil {
// 		cart.User = newCart.User
// 	}

// }

// func (cart *Cart) GetUpdate() {
// }

// func (cart *Cart) GetUUID() uuid.UUID {
// 	return cart.id
// }
