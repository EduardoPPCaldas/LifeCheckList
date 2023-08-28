package repository

import (
	"github.com/EduardoPPCaldas/LifeCheckList/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(uuid.UUID) (*entity.User, error)
	Create(entity.User) (*entity.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(database *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: database,
	}
}

func (repo GormUserRepository) GetById(id uuid.UUID) (*entity.User, error) {
	var user entity.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo GormUserRepository) Create(user entity.User) (*entity.User, error) {
	result := repo.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}