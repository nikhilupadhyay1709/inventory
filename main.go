package main

import (
	"github.com/dwahyudi/inventory/configs/database"
	"github.com/dwahyudi/inventory/configs/routes"
	"log"
)

func main() {
	database.Prepare()
	web := routes.GenerateRoutes()

	log.Fatal(web.Run())
}
