package service

import "part2/internal/model"

type UserStorage struct {
	db map[string]*model.User
}

type UserStorageUpdater struct {
	storage UserStorage
}

func (u *UserStorageUpdater) Update(user *model.User) {
	u.storage.db[user.Login] = user
}
