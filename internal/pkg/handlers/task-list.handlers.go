package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (gh *GlobalHandler) createOneList(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskListData := new(models.TaskList)
	if err := ctx.ShouldBindJSON(taskListData); err != nil {
		err = isEmptyBodyErr(err)
		gh.callResponseGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := gh.services.TaskListServiceCreate(personId, taskListData)
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"list_id": listId,
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
