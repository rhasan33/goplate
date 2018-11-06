package models

import (
	"github.com/jinzhu/gorm"
)

// UserSettings model for user account
type UserSettings struct {
	gorm.Model
	Username     string `gorm:"not null;unique_index" json:"username"`
	Email        string `gorm:"not null;unique_index" json:"email"`
	Password     string `gorm:"not null" json:"password"`
	ReferralCode string `gorm:"not null;unique" json:"referral_code"`
}
