package adaptor

import (
	"errors"
	"part2/internal/model"
	"sync"

	"github.com/google/uuid"
)

type ConcurrentUserStorage struct {
	db map[string]*model.User
	mx sync.Mutex
}

func CreateConcurrentUserStorage() *ConcurrentUserStorage {
	return &ConcurrentUserStorage{
		db: make(map[string]*model.User),
		mx: sync.Mutex{},
	}
}

func validateUser(user *model.User) error {
	if user.ID == uuid.Nil {
		return errors.New("ID is null")
	}
	if user.Login == "" {
		return errors.New("Login is null")
	}
	if user.Password == "" {
		return errors.New("Login is null")
	}
	return nil

}

func (u *ConcurrentUserStorage) Get(login string) *model.User {
	if u.db[login] != nil {
		return u.db[login]
	}
	return nil
}

type UserStorageUpdater struct {
	storage ConcurrentUserStorage
}

func (u *UserStorageUpdater) Update(user *model.User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}
	if u.storage.db[user.Login] == nil {
		return errors.New("This login doesn't exist")
	}
	u.storage.db[user.Login] = user
	return nil
}

func (u *UserStorageUpdater) Add(user *model.User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}
	if u.storage.db[user.Login] != nil {
		return errors.New("This login already exist")
	}
	u.storage.db[user.Login] = user
	return nil
}
