package main

import (
	"os"

	"github.com/whitejokeer/clevertitest/api"
	datamanager "github.com/whitejokeer/clevertitest/database"
)

func main() {
	port := os.Getenv("PORT")
	dbConfig := datamanager.GetConfig()

	app := &api.App{}
	app.Initialize(dbConfig)
	app.Run(":" + port)
}
