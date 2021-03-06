package handlers

import (
	"github.com/gin-gonic/gin"
)

type responseGenericError struct {
	Message string `json:"message"`
}

func (gh *GlobalHandler) callResponseGenericError(ctx *gin.Context, statusCode int, message string) {
	gh.log.Error(message)
	ctx.AbortWithStatusJSON(statusCode, &responseGenericError{Message: message})
}
