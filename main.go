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

	api := api.NewLifeCheckListApi()

	_, err := database.SetDatabase()

	if err != nil {
		log.Fatal(err)
	}

	api.Run()
}