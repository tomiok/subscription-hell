package subscription

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nick     string `gorm:"unique"`
	Password string
}

type Storage interface {
	Create(nick, password string) (User, error)
	Login(nick, password string) (string, error)
}

type Service struct {
	Secret string
}
