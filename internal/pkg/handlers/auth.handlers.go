package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (gh *GlobalHandler) signUp(ctx *gin.Context) {
	person := new(models.Person)

	err := ctx.BindJSON(person)
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := gh.services.AuthServiceCreatePerson(person)
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (gh *GlobalHandler) login(ctx *gin.Context) {
	//
}
