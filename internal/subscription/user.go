package subscription

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nick     string `gorm:"unique"`
	Password string
	Subs     []Sub
}

type UserStorage interface {
	Create(nick, password string) (User, error)
	Login(nick, password string) (string, error)
}

type UserService struct {
	Secret string
}
