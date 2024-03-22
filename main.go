package main

import (
	"mygram/database"
	"mygram/router"
)

func main() {
    database.StartDB()
    r := router.SetupRouter()
    r.Run(":8080")
}
