package main

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/db_init"
	"NSK_mayoralty_app/router"
	"fmt"
	"log"
	"os"
)

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
