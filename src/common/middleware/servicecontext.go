package middleware

import (
	"github.com/gin-gonic/gin"
	"lookingforpartner/common"
)

func InjectSvcCtxMiddleware(svcCtx interface{}) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set(common.SvcCtx, svcCtx)
		c.Next()
	}
}
