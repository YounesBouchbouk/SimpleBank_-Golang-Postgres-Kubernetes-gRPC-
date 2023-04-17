package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) != 2 {
			err := errors.New("invalide authorization formate ")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		if strings.ToLower(fields[0]) != authorizationTypeBearer {
			err := errors.New("unsuported authorization type")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accesToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accesToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)

		ctx.Next()
	}

}
