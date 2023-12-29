package subscription

import "gorm.io/gorm"

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
}
