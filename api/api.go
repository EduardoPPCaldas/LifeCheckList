package api

import "github.com/labstack/echo/v4"

type LifeCheckListApi struct {}

func NewLifeCheckListApi() *LifeCheckListApi {
	return &LifeCheckListApi{}
}

func (l *LifeCheckListApi) Run() {
	e := echo.New()
	
	e.Logger.Fatal(e.Start(":1323"))
}