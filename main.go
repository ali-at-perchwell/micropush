package main

import (
	"micropush/service"
	_ "net/http/pprof"
	"os"
)

func main() {
	a := service.App{}
	// TODO: make this into a conifg
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_SSLMODE"))

	a.Run()
}
