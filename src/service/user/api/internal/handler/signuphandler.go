package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"lookingforpartner/common"
	httpUtil2 "lookingforpartner/common/httputils"
	"lookingforpartner/idl/pb/user"
	"lookingforpartner/service/user/api/internal/svc"
	"net/http"
)

func SingUpHandler(c *gin.Context) {
	var req user.UserSignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			httpUtil2.ResponseError(c, common.CodeInvalidParam, http.StatusBadRequest)
		} else {
			errMsg := httpUtil2.RemoveTopStruct(errors.Translate(httpUtil2.Trans))
			httpUtil2.ResponseErrorWithMsg(c, common.CodeInvalidParam, http.StatusBadRequest, errMsg)
		}
		return
	}

	fmt.Println("signup api")

	svcCtx := svc.GetServiceContext(c)
	userClient := svcCtx.UserClient

	_, err := userClient.UserSignup(c, &req)
	if err != nil {
		zap.L().Error("Invoke signup RPC failed", zap.Error(err))
		if errors.Is(err, common.ErrorUserExist) {
			httpUtil2.ResponseError(c, common.CodeUserExists, http.StatusBadRequest)
		} else {
			httpUtil2.ResponseError(c, common.CodeServerBusy, http.StatusInternalServerError)
		}
		return
	}

	httpUtil2.ResponseSuccess(c, nil)
}
