package httpUtil

import (
	"github.com/gin-gonic/gin"
	"lookingforpartner/common"
)

const ContextUserIDKey = "userID"

const (
	DefaultPostPageValue = 1
	DefaultPostSizeValue = 10

	DefaultCommentPageValue = 1
	DefaultCommentSizeValue = 15
)

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

func GetCurrentUser(c *gin.Context) (int64, error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		return 0, common.ErrorUserNotLogin
	}
	userID, ok := uid.(int64)
	if !ok {
		return 0, common.ErrorUserNotLogin
	}
	return userID, nil
}
