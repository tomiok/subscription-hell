package subscription

import (
	"gorm.io/gorm"
)

type (
	SubType string
	Freq    string
)

type Sub struct {
	gorm.Model
	Name      string
	ManagedIn string
	PaidWith  string
	Type      SubType
	Frequency Freq
	UserID    uint
}

type SubRequest struct {
	Name      string  `json:"name"`
	ManagedIn string  `json:"managed_in"`
	PaidWith  string  `json:"paid_with"`
	Type      SubType `json:"type"`
	Frequency Freq    `json:"frequency"`
	UserID    uint    `json:"-"`
}

type Storage interface {
	Add(s Sub) (Sub, error)
	ViewAll(userID uint) ([]Sub, error)
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) AddSub(req SubRequest) (Sub, error) {
	var sub = Sub{
		Name:      req.Name,
		ManagedIn: req.ManagedIn,
		PaidWith:  req.PaidWith,
		Type:      req.Type,
		Frequency: req.Frequency,
		UserID:    req.UserID,
	}

	return s.storage.Add(sub)
}
