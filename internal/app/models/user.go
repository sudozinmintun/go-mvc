package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func (u *User) SetPassword(p string) error {
	h, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		return err
	}
	u.Password = string(h)
	return nil
}

func (u *User) CheckPassword(p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}

func SeedAdmin(db *gorm.DB) {
	admin := User{Email: "admin@example.com"}
	admin.SetPassword("secret123")
	db.Create(&admin)
}
