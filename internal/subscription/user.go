package subscription

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Nick     string `gorm:"unique"`
	Password string
	Subs     []Sub
}

type UserStorage interface {
	Create(user User) (User, error)
	Login(nick string) (User, error)
}

type UserService struct {
	storage UserStorage
	Secret  string
}

func NewUserService(secret string, storage UserStorage) *UserService {
	return &UserService{
		storage: storage,
		Secret:  secret,
	}
}

func (us *UserService) CreateUser(nick, pass string) (User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	var user = User{
		Nick:     nick,
		Password: string(password),
		Subs:     nil,
	}

	return us.storage.Create(user)
}

func (us *UserService) LoginUser(nick, pass string) (User, string, error) {
	user, err := us.storage.Login(nick)
	if err != nil {
		return User{}, "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return User{}, "", err
	}

	return user, signToken(user.ID, us.Secret), nil
}

func signToken(id uint, key string) string {
	claims := make(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * 172999).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedString, _ := token.SignedString([]byte(key))

	return signedString
}
