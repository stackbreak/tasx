package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (gh *GlobalHandler) createOneTask(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, ErrInvalidListIdParam.Error())
		return
	}

	inputData := new(models.Task)
	if err := ctx.ShouldBindJSON(inputData); err != nil {
		err = isEmptyBodyErr(err)
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := gh.services.TaskServiceCreateOne(personId, listId, inputData)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"id": id,
	})
}

func (gh *GlobalHandler) getAllTasks(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, ErrInvalidListIdParam.Error())
		return
	}

	tasks, err := gh.services.TaskServiceGetAll(personId, listId)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"tasks": tasks,
	})
}

func (gh *GlobalHandler) getOneTaskById(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, ErrInvalidTaskIdParam.Error())
		return
	}

	task, err := gh.services.TaskServiceGetOne(personId, taskId)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"task": task,
	})
}

func (gh *GlobalHandler) updateOneTask(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, ErrInvalidTaskIdParam.Error())
		return
	}

	inputData := new(models.InputUpdateTask)
	if err := ctx.ShouldBindJSON(inputData); err != nil {
		err = isEmptyBodyErr(err)
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = gh.services.TaskServiceUpdateOne(personId, taskId, inputData)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &respStatus{"ok"})
}

func (gh *GlobalHandler) deleteOneTask(ctx *gin.Context) {
	personId, err := extractPersonIdFromCtx(ctx)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusBadRequest, ErrInvalidTaskIdParam.Error())
		return
	}

	err = gh.services.TaskServiceDeleteOne(personId, taskId)
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &respStatus{"ok"})
}
