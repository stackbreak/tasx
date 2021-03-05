package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	CtxUserIdKey        = "userId"
)

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

	ctx.Set(CtxUserIdKey, userId)
}
