package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/whitejokeer/clevertitest/api"
)

func main() {
	// TODO: Remove this line when you're ready to implement the code using Docker
	godotenv.Load()

	port := os.Getenv("PORT")

	app := &api.App{}
	app.Initialize()
	app.Run(":" + port)
}
