package main

import (
	mysqldb "country-dropdown-gin/connection"
	"country-dropdown-gin/routes"
	"log"
)

func main() {
	log.Println("listening on http://localhost:8040")
	db := mysqldb.SetupDB()
	r := routes.SetupRoutes(db)
	r.Run(":8040")

}
