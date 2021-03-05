package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gh *GlobalHandler) createOneList(ctx *gin.Context) {
	userId, _ := ctx.Get(CtxUserIdKey)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}

func (gh *GlobalHandler) getAllLists(ctx *gin.Context) {
	//
}

func (gh *GlobalHandler) getOneListById(ctx *gin.Context) {
	//
}

func (gh *GlobalHandler) updateOneList(ctx *gin.Context) {
	//
}

func (gh *GlobalHandler) deleteOneList(ctx *gin.Context) {
	//
}
