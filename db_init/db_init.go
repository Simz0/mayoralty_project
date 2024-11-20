package db_init

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/models"
	"log"
)

func InitDataBases() bool {
	cityPointsErr := database.DB.AutoMigrate(&models.City_points{})
	if cityPointsErr != nil {
		log.Fatal("City_points migration error")
	}
	depErr := database.DB.AutoMigrate(&models.Departament{})
	if depErr != nil {
		log.Fatal("Departament migration error")
	}
	districtErr := database.DB.AutoMigrate(&models.District{})
	if districtErr != nil {
		log.Fatal("District migration error")
	}
	sessionsErr := database.DB.AutoMigrate(&models.Sessions{})
	if sessionsErr != nil {
		log.Fatal("Sessions migration error")
	}
	userErr := database.DB.AutoMigrate(&models.User{})
	if userErr != nil {
		log.Fatal("User migration error")
	}
	workPointsErr := database.DB.AutoMigrate(&models.Work_points{})
	if workPointsErr != nil {
		log.Fatal("Work_points migration error")
	}

	return true
}
