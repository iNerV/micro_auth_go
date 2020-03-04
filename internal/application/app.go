package application

import (
	"github.com/gorilla/mux"
	"log"
	"micro_auth/core"
	"micro_auth/internal/database"
	"micro_auth/internal/router"
	"net/http"
	"sync"
)

type App struct {
	Router *mux.Router
}

var (
	instance *App
	once     sync.Once
)

func GetApp() *App {
	once.Do(func() {
		instance = &App{}
		instance.Initialize()
	})
	return instance
}

func (app *App) Initialize() {
	app.initializeRoutes()
	initializeDB()
}

func (app *App) Run() {
	port := core.AppConfig.Port
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}

func (app *App) initializeRoutes() {
	app.Router = mux.NewRouter().StrictSlash(true)
	router.BuildRouter(app.Router)
}

func initializeDB() {
	dsn := core.AppConfig.DatabaseURL
	database.Open(dsn)
}
