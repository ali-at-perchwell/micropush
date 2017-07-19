package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	web_push "/micropush/resource/web/push" // the naming on this might be in violation
	web_subscription "/micropush/resource/web/subscription"

	devicepush "/micropush/resource/device/push"
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

	a.Router = gin.Default() //mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run() { /// put in addr
	log.Fatal(a.Router.Run()) //http.ListenAndServe(":8080", a.Router))
}

func HealthCheck(c *gin.Context) {
	c.String(200, "Success")
}
func (a *App) initializeRoutes() {
	a.Router.GET("/health", HealthCheck)

	a.Router.POST("/subscribe/web", web_subscription.Create)
	a.Router.DELETE("/subscribe/web", web_subscription.Delete)
	// a.Router.DELETE("/subscribe/web/:id", web_subscription.Delete)  so here we would inspect the route params rather than json
	a.Router.PUT("/subscribe/web", web_subscription.Update)

	a.Router.POST("/subscribe/device", web_subscription.Create)
	a.Router.DELETE("/subscribe/device", web_subscription.Delete)
	a.Router.PUT("/subscribe/device", web_subscription.Update)

	a.Router.POST("/subscribe/web/push", web_push.Create)
	a.Router.PUT("/subscribe/web/push", web_push.Create) // we can have an edit, no? but no delete to re-alert

	a.Router.POST("/subscribe/device/push", device_push.Create)
}
