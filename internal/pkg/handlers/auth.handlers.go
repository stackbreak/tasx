package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/stackbreak/tasx/internal/pkg/models"
)

func (gh *GlobalHandler) signUp(ctx *gin.Context) {
	personData := new(models.Person)

	err := ctx.ShouldBindJSON(personData)
	if err != nil {
		err = isEmptyBodyErr(err)
		gh.callResponseGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := gh.services.AuthServiceCreatePerson(personData)
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (gh *GlobalHandler) login(ctx *gin.Context) {
	loginData := new(loginInput)

	err := ctx.ShouldBindJSON(loginData)
	if err != nil {
		err = isEmptyBodyErr(err)
		gh.callResponseGenericError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := gh.services.AuthServiceGenerateToken(loginData.Username, loginData.Password)
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
