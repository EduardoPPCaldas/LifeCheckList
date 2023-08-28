package main

import "github.com/EduardoPPCaldas/LifeCheckList/api"

func main() {
	api := api.NewLifeCheckListApi()
	api.Run()
}