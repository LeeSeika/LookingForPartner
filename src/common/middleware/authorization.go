package middleware

import (
	"github.com/gin-gonic/gin"
	"lookingforpartner/common"
	httpUtil "lookingforpartner/pkg/httputils"
	"lookingforpartner/pkg/jwtx"
	"net/http"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			httpUtil.ResponseError(c, common.CodeAuthNotFound, http.StatusBadRequest)
			c.Abort()
			return
		}

		mc, err := jwtx.ParseToken(authHeader)
		if err != nil {
			// controller.ResponseError(c, biz.CodeInvalidToken)
			httpUtil.ResponseError(c, common.CodeInvalidToken, http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set(httpUtil.ContextUserIDKey, mc.UserID)
		c.Next()
	}
}
