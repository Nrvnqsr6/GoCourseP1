package action

import (
	"part1/internal/consts"
	"part1/internal/model"
)

type CartCreateWithStatReport struct {
	StatusChannel chan model.Status
}

func (cartCreate *CartCreateWithStatReport) Do(user *model.User) *model.Cart {
	cart := model.CreateCart(user)
	status := model.CreateStatus(cart, consts.CREATED)
	go func() {
		cartCreate.StatusChannel <- *status
	}()
	return cart
}

type CartUpdateWithStatReport struct {
	StatusChannel chan model.Status
	Cart          *model.Cart
}

func (cartUpdate *CartUpdateWithStatReport) Do(newCart *model.Cart) {
	cartUpdate.Cart.Update(newCart)
	status := model.CreateStatus(cartUpdate.Cart, consts.CHANGED)
	go func() {
		cartUpdate.StatusChannel <- *status
	}()
}
