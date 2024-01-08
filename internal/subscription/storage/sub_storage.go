package storage

import (
	"github.com/tomiok/subscription-hell/internal/subscription"
	"github.com/tomiok/subscription-hell/pkg/database"
)

type SubStorage struct{}

func NewSubStorage() *SubStorage {
	return &SubStorage{}
}

func (s *SubStorage) Add(sub subscription.Sub) (subscription.Sub, error) {
	if err := database.DB.Save(&sub).Error; err != nil {
		return subscription.Sub{}, err
	}

	return sub, nil
}

func (s *SubStorage) ViewAll(userID uint) ([]subscription.Sub, error) {
	var res []subscription.Sub
	if err := database.DB.Find(&res, "user_id = $1", userID).Error; err != nil {
		return nil, err
	}

	return res, nil
}
