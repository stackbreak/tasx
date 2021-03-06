package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	CtxPersonIdKey      = "personId"
)

var ErrPersonIdNotExtracted = errors.New("pkg.handlers: unable to extract person_id")

func (gh *GlobalHandler) mdwUserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		gh.callResponseGenericError(ctx, http.StatusUnauthorized, "empty \"Authorization\" header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		gh.callResponseGenericError(ctx, http.StatusUnauthorized, "invalid \"Authorization\" header")
		return
	}

	userId, err := gh.services.AuthServiceParseToken(headerParts[1])
	if err != nil {
		gh.callResponseGenericError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(CtxPersonIdKey, userId)
}

func extractPersonIdFromCtx(ctx *gin.Context) (int, error) {
	personId, ok := ctx.Get(CtxPersonIdKey)
	if !ok {
		return -1, ErrPersonIdNotExtracted
	}

	personIdInt, ok := personId.(int)
	if !ok {
		return -1, ErrPersonIdNotExtracted
	}

	return personIdInt, nil
}
