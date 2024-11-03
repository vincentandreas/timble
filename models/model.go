package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	Gender string `gorm:"size:255;unique;not null" json:"gender"`
	Quota uint `gorm:"not null" json:"quota"`
	QuotaLatestUpdate time.Time `gorm:"not null"`
}

type UserPreference struct {
	gorm.Model
	UserID uint64 `gorm:"primarykey" json:"-"`
	InterestGender string `gorm:"not null" json:"interest_gender"`
}

type History struct {
	gorm.Model
	ID          uint64 `gorm:"primarykey" json:"-"`
	UserID uint64 `gorm:"primarykey" json:"user_id"`
	WatchedUser User `gorm:"foreignkey:WatchedUserID"`
	WatchedUserID uint64 `gorm:"not null" json:"watched_user_id"`
	FeelMatch bool `gorm:"primarykey" json:"feel_match"`	
}

type UserLogin struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}



type Feeling struct{
	WatchedUserID uint64 ` json:"watched_user_id"`
	FeelMatch bool ` json:"feel_match"`	
}