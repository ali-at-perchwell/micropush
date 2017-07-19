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

	"/micropush/resource/push"
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

func (a *App) initializeRoutes() {
	a.Router.GET("/healthCheck", pushResource.HealthCheck)

	a.Router.POST("/token/token-web/:id", pushResource.CreateWebToken)
	a.Router.DELETE("/token/token-web/:id", pushResource.DeleteWebToken)
	a.Router.POST("/token/token-device/:id", pushResource.CreateDeviceToken)
	a.Router.DELETE("/token/token-device/:id", pushResource.DeleteDeviceToken)

	a.Router.POST("/initialize-push", pushResource.InitializePush)
}

func (a *App) getSubscriptions(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
}
