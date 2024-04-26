package middleware

import (
	"github.com/gin-gonic/gin"
	"lookingforpartner/common"
	httpUtil2 "lookingforpartner/common/httputils"
	"lookingforpartner/common/jwtx"
	"net/http"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			httpUtil2.ResponseError(c, common.CodeAuthNotFound, http.StatusBadRequest)
			c.Abort()
			return
		}

		mc, err := jwtx.ParseToken(authHeader)
		if err != nil {
			// controller.ResponseError(c, biz.CodeInvalidToken)
			httpUtil2.ResponseError(c, common.CodeInvalidToken, http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set(httpUtil2.ContextUserIDKey, mc.UserID)
		c.Next()
	}
}
