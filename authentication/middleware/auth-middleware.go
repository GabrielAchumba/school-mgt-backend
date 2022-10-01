package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GabrielAchumba/school-mgt-backend/authentication/rest"
	"github.com/GabrielAchumba/school-mgt-backend/authentication/token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorizationPayload"
)

func GetAuthorizationPayload(ctx *gin.Context) (token.Payload, bool) {
	payload, exist := ctx.Get(AuthorizationPayloadKey)
	if exist == false {
		return token.Payload{}, exist
	}
	return payload.(token.Payload), exist
}

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, rest.GetError(http.StatusUnauthorized, "Authorization header is not provided"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, rest.GetError(http.StatusUnauthorized, "Invalid authorization header format"))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("Unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, rest.GetError(http.StatusUnauthorized, err.Error()))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, rest.GetError(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.Set(AuthorizationPayloadKey, *payload)
		ctx.Next()
	}
}
