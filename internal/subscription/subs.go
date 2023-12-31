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

type Storage interface {
	Add(ctx context.Context, s Sub) (Sub, error)
	ViewAll(ctx context.Context) ([]Sub, error)
}
