package database

import (
	"github.com/EduardoPPCaldas/LifeCheckList/domain/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetDatabase() (*gorm.DB, error) {
	dsn := viper.Get("DATABASE_DNS").(string)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{})

	return db, nil
}