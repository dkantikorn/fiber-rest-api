package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"primary_key;auto_increment;" json:"id"`
	Username    string    `validate:"required,min=5,max=50" gorm:"size:50;not null;unique;" json:"username"`
	Password    string    `validate:"required,min=32,max=255" gorm:"size:255;not null;" json:"password"`
	FirstName   string    `validate:"required,min=3,max=255" json:"first_name"`
	LastName    string    `validate:"required,min=3,max=255" json:"last_name"`
	Email       string    `validate:"required,min=3,max=255,email" gorm:"size:255;not null;unique;" json:"email"`
	PicturePath string    `gorm:"size:512;" json:"picture_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
