package main

import (
	"novocaine-dev/database"
	"novocaine-dev/routes"
)

func main() {
	db := database.SetupDatabase()
	r := routes.SetupRoutes(db)
	r.Run()
}
