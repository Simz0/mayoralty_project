package main

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/db_init"
	_ "NSK_mayoralty_app/docs"
	"NSK_mayoralty_app/router"
	"fmt"
	"log"
	"os"
)

// @title Проект для мэрии Новосибирска по отслеживанию и визуализации геоточек и маршрутов
// @version 1.0
// description

// @host undefined
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name NSK_Mayoralty

func main() {
	database.ConnectDB()

	result := db_init.InitDataBases()
	if result {
		fmt.Println("Connect to BD is succefull")
	} else {
		log.Fatal("Error with connect to database")
	}
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
