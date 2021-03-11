package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/stackbreak/tasx/internal/pkg/services"
)

type GlobalHandler struct {
	services *services.Services
	log      *logrus.Logger
}

func NewGlobalHandler(services *services.Services, log *logrus.Logger) *GlobalHandler {
	return &GlobalHandler{services, log}
}

func (gh *GlobalHandler) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &gin.H{
		"status": "available",
	})
}

func (gh *GlobalHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/healthz", gh.healthCheck)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", gh.signUp)
		auth.POST("/login", gh.login)
	}

	api := router.Group("/api", gh.mdwUserIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", gh.createOneList)
			lists.GET("/", gh.getAllLists)
			lists.GET("/:id", gh.getOneListById)
			lists.PUT("/:id", gh.updateOneList)
			lists.DELETE("/:id", gh.deleteOneList)

			items := lists.Group(":id/items")
			{
				items.POST("/", gh.createOneTask)
				items.GET("/", gh.getAllTasks)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", gh.getOneTaskById)
			items.PUT("/:id", gh.updateOneTask)
			items.DELETE("/:id", gh.deleteOneTask)
		}
	}

	return router
}
