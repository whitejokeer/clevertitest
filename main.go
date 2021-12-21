package main

import (
	"os"

	"github.com/whitejokeer/clevertitest/api"
)

func main() {
	port := os.Getenv("PORT")

	app := &api.App{}
	app.Initialize()
	app.Run(":" + port)
}
