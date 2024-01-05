package subscription

import (
	"context"
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
	UserID    uint    `json:"user_id"`
}

type Storage interface {
	Add(ctx context.Context, s Sub) (Sub, error)
	ViewAll(ctx context.Context) ([]Sub, error)
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{
		storage: s,
	}
}

func (s *Service) AddSub() {

}
