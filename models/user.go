package models

import (
	"time"
)

type User struct {
	ID              uint   `gorm:"primary_key" json:"id"`
	Email           string `gorm:"type:varchar(255);unique_index;not null" json:"email"`
	Username        string `gorm:"type:varchar(255);unique_index;not null" json:"username"`
	Password        string `gorm:"type:varchar(255);not null" json:"password"`
	Age             int    `gorm:"size:11;not null" json:"age"`
	ProfileImageURL string `gorm:"type:varchar(255);not null" json:"profile_image_url"`

	// User Has Many Photos (One to Many)
	// User dapat memposting banyak Photos
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User Has Many SocialMedias (One to Many)
	// User dapat punya banyak akun di SocialMedias
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User Has Many Comments
	// User dapat punya banyak Comments di tiap photos
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
