package di

import (
	"github.com/EduardoPPCaldas/LifeCheckList/api/controller"
	"github.com/EduardoPPCaldas/LifeCheckList/application/usecase"
	"github.com/EduardoPPCaldas/LifeCheckList/infrastructure/repository"
	"gorm.io/gorm"
)

func InjectUserController(db *gorm.DB) *controller.UserController {
	repo := repository.NewGormUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	return controller.NewUserController(uc)
}