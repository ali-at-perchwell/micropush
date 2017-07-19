package main

import (
	"micropush/service"
	"os"
)

func main() {
	a := service.App{}
	a.Initialize( // TODO: make this into a conifg
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run()
}
