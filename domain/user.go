package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `gorm:"not null"`
	Email          string `gorm:"index:idx_users_email,unique;not null"`
	HashedPassword string `gorm:"not null"`
	IsVerified     bool   `gorm:"not null"`
}

type UserUsecase interface {
	Register(fullname, email, plainPassword string) error
	Login(email, plainPassword string) (User, Wallet, string, error)
}

type UserRepository interface {
	Insert(fullname, email, hashedPassword string, isVerified bool) error
	GetByEmail(email string) (User, Wallet, error)
}
