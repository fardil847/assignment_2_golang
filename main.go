package main

import (
	"assignment_2_golang/config"
	"assignment_2_golang/controllers"
	"assignment_2_golang/routers"
)

var PORT = ":8080"

func main() {
	db := config.StartDB()
	controller := controllers.New(db)

	routers.StartServer(controller).Run(PORT)
}
