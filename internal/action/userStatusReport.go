package action

import (
	"part1/internal/consts"
	"part1/internal/model"
)

type UserCreateWithStatReport struct {
	StatusChannel chan model.Status
}

func (userCreate *UserCreateWithStatReport) Do(name string) *model.User {
	user := model.CreateUser(name)
	status := model.CreateStatus(user, consts.CREATED)
	go func() {
		userCreate.StatusChannel <- *status
	}()
	return user
}

type UserUpdateWithStatReport struct {
	StatusChannel chan model.Status
	User          *model.User
}

func (userUpdate *UserUpdateWithStatReport) Do(newName string) {
	userUpdate.User.Update(newName)
	status := model.CreateStatus(userUpdate.User, consts.CHANGED)
	go func() {
		userUpdate.StatusChannel <- *status
	}()
}
