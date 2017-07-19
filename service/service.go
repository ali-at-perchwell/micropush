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
func (a *App) Initialize(user, password, dbname string) { // call this init() ?
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = gin.Default()
	a.initializeRoutes()
}

func (a *App) Run() { /// put in addr
	log.Fatal(a.Router.Run()) //http.ListenAndServe(":8080", a.Router))
}

func HealthCheck(c *gin.Context) {
	c.String(200, "Success") // or c.JSON(200,
}
func (a *App) initializeRoutes() {
	a.Router.GET("/health", HealthCheck)

	a.Router.POST("/subscribe/web", web.CreateSubscription)
	a.Router.DELETE("/subscribe/web", web.DeleteSubscription)

	a.Router.PUT("/subscribe/web", web.UpdateSubscription)

	a.Router.POST("/subscribe/device", device.CreateSubscription)
	a.Router.DELETE("/subscribe/device", device.DeleteSubscription)
	a.Router.PUT("/subscribe/device", device.UpdateSubscription)

	a.Router.POST("/push/web", web.CreatePush)
	a.Router.PUT("/push/web", web.UpdatePush)

	a.Router.POST("/push/device", device.CreatePush)
}
