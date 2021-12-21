package main

import (
	"github.com/whitejokeer/clevertitest/api"
)

func main() {
	port := "8080" // os.Getenv("PORT")

	app := &api.App{}
	app.Initialize()
	app.Run(":" + port)
}
