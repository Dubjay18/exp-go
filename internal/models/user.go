package models

import (
	"exp-go/internal/database/postgresql"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	gorm.Model
}

// HashPassword hashes the user's password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword checks if the provided password matches the stored password
func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}

func (user *User) Create(db *gorm.DB) error {
	err := postgresql.CreateOneRecord(db, user)
	if err != nil {
		return err
	}
	return nil
}
