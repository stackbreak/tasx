package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/stackbreak/tasx/internal/pkg/service"
)

type GlobalHandler struct {
	services *service.Service
}

func NewGlobalHandler(services *service.Service) *GlobalHandler {
	return &GlobalHandler{services}
}

func (gh *GlobalHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	auth.POST("/signup", gh.signUp)
	auth.POST("/login", gh.login)

	api := router.Group("/api")

	lists := api.Group("/lists")
	lists.POST("/", gh.createOneList)
	lists.GET("/", gh.getAllLists)
	lists.GET("/:id", gh.getOneListById)
	lists.PUT("/:id", gh.updateOneList)
	lists.DELETE("/:id", gh.deleteOneList)

	items := lists.Group(":id/items")
	items.POST("/", gh.createOneItem)
	items.GET("/", gh.getAllItems)
	items.GET("/:item_id", gh.getOneItemById)
	items.PUT("/:item_id", gh.updateOneItem)
	items.DELETE("/:item_id", gh.deleteOneItem)

	return router
}
