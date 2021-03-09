package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	CtxPersonIdKey      = "personId"
)

func (gh *GlobalHandler) mdwUserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		gh.callRespGenericError(ctx, http.StatusUnauthorized, "empty \"Authorization\" header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		gh.callRespGenericError(ctx, http.StatusUnauthorized, "invalid \"Authorization\" header")
		return
	}

	userId, err := gh.services.AuthServiceParseToken(headerParts[1])
	if err != nil {
		gh.callRespGenericError(ctx, http.StatusUnauthorized, err.Error())
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
