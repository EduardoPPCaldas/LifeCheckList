package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string    `gorm:"uniqueIndex"`
	Name     string
	Password string
}

func NewUser(email string, name string, password string) *User {
	return &User{
		Email: email,
		Name: name,
		Password: password,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}
