package api

import (
	"github.com/EduardoPPCaldas/LifeCheckList/infrastructure/di"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LifeCheckListApi struct {
	db *gorm.DB
}

func NewLifeCheckListApi(db *gorm.DB) *LifeCheckListApi {
	return &LifeCheckListApi{
		db: db,
	}
}

func (l *LifeCheckListApi) Run() {
	e := echo.New()

	userController := di.InjectUserController(l.db);

	e.POST("/signup", userController.SignUp)
	
	e.Logger.Fatal(e.Start(":1323"))
}