package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

type respStatus struct {
	Status string `json:"status"`
}

type respGenericError struct {
	Message string `json:"message"`
}

type respAllTaskLists struct {
	Data []models.TaskList `json:"data"`
}

func (gh *GlobalHandler) callRespGenericError(ctx *gin.Context, statusCode int, message string) {
	gh.log.Error(message)
	ctx.AbortWithStatusJSON(statusCode, &respGenericError{Message: message})
}
