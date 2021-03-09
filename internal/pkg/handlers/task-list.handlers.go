package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (gh *GlobalHandler) createOneList(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskListData := new(models.TaskList)
	if err := ctx.ShouldBindJSON(taskListData); err != nil {
		err = isEmptyBodyErr(err)
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := gh.services.TaskListServiceCreate(personId, taskListData)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"list_id": listId,
	})
}

func (gh *GlobalHandler) getAllLists(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := gh.services.TaskListServiceGetAll(personId)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &respAllTaskLists{
		Data: lists,
	})
}

func (gh *GlobalHandler) getOneListById(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskListId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	list, err := gh.services.TaskListServiceGetOne(personId, taskListId)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (gh *GlobalHandler) updateOneList(ctx *gin.Context) {
	//
}

func (gh *GlobalHandler) deleteOneList(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskListId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = gh.services.TaskListServiceDeleteOne(personId, taskListId)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &respStatus{"ok"})
}
