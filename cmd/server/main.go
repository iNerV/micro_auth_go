package main

import (
	"micro_auth/internal/application"
)

func main() {
	app := application.App{}
	app.Initialize()

	app.Run()
}
