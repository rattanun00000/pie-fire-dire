package app

import (
	"pie-fire-dire/internal/handler"
	"pie-fire-dire/internal/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	router       *gin.Engine
	beefService  *service.BeefService
	baconService *service.MeatIpsumService
	beefHandler  *handler.BeefHandler
}

func NewApp() *App {
	gin.SetMode(gin.ReleaseMode)

	baconService := service.NewMeatIpsumService()
	beefService := service.NewBeefService()

	beefHandler := handler.NewBeefHandler(beefService, baconService)

	router := gin.Default()

	return &App{
		router:       router,
		beefService:  beefService,
		baconService: baconService,
		beefHandler:  beefHandler,
	}
}

func (a *App) Run(addr string) error {
	a.setupRoutes()

	return a.router.Run(addr)
}

func (a *App) setupRoutes() {
	a.router.GET("/beef/summary", a.beefHandler.GetBeefSummary)
}
