package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq" // handles bad connections for db/sql

	"micropush/device"
	"micropush/web"
)

// App represents the application
type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

// Initialize sets up the database connection and routes for the app
func (a *App) Initialize(user, password, dbname, sslmode string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = gin.Default()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(a.Router.Run(addr))
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, "Success")
}

func (a *App) initializeRoutes() {
	// Health Check
	a.Router.GET("/health", HealthCheck)

	// Web Subscription
	a.Router.POST("/subscribe/web", web.CreateSubscription)
	a.Router.DELETE("/subscribe/web", web.DeleteSubscription)
	a.Router.PUT("/subscribe/web", web.UpdateSubscription)

	// Web Push
	a.Router.POST("/push/web", web.CreatePush)
	a.Router.PUT("/push/web", web.UpdatePush)

	// Device Subscription
	a.Router.POST("/subscribe/device", device.CreateSubscription)
	a.Router.DELETE("/subscribe/device", device.DeleteSubscription)
	a.Router.PUT("/subscribe/device", device.UpdateSubscription)

	// Device Push
	a.Router.POST("/push/device", device.CreatePush)
}
