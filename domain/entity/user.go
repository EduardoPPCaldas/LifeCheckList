package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string
	Email    string
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}
