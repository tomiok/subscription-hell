package storage

import (
	"github.com/tomiok/subscription-hell/internal/subscription"
)

type UserStorage struct {
}

func (u UserStorage) Create(nick, password string) (subscription.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserStorage) Login(nick, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}
