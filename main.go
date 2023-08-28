package main

import (
	"log"

	"github.com/EduardoPPCaldas/LifeCheckList/api"
	"github.com/EduardoPPCaldas/LifeCheckList/infrastructure/database"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db, err := database.SetDatabase()
	
	if err != nil {
		log.Fatal(err)
	}
	
	api := api.NewLifeCheckListApi(db)

	api.Run()
}