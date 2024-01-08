package storage

import (
	"github.com/tomiok/subscription-hell/internal/subscription"
	"github.com/tomiok/subscription-hell/pkg/database"
)

type UserStorage struct{}

func NewUserStorage() *UserStorage {
	return &UserStorage{}
}

func (u UserStorage) Create(user subscription.User) (subscription.User, error) {
	if err := database.DB.Save(&user).Error; err != nil {
		return subscription.User{}, err
	}

	return user, nil
}

func (u UserStorage) Login(nick string) (subscription.User, error) {
	var user subscription.User
	if err := database.DB.Find(&user, "nick=$1", nick).Error; err != nil {
		return subscription.User{}, err
	}

	return user, nil
}
