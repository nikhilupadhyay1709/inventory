package main

import (
	"inventory/configs/database"
	"inventory/configs/routes"
	"log"
)

func main() {
	database.Prepare()
	web := routes.GenerateRoutes()

	log.Fatal(web.Run())
}
