package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"lookingforpartner/common"
	"lookingforpartner/idl/pb/user"
	httpUtil "lookingforpartner/pkg/httputils"
	"lookingforpartner/service/user/api/internal/svc"
	"net/http"
)

func SingUpHandler(c *gin.Context) {
	var req user.UserSignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("Sign up with invalid param", zap.Error(err))
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			httpUtil.ResponseError(c, common.CodeInvalidParam, http.StatusBadRequest)
		} else {
			errMsg := httpUtil.RemoveTopStruct(errors.Translate(httpUtil.Trans))
			httpUtil.ResponseErrorWithMsg(c, common.CodeInvalidParam, http.StatusBadRequest, errMsg)
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
			httpUtil.ResponseError(c, common.CodeUserExists, http.StatusBadRequest)
		} else {
			httpUtil.ResponseError(c, common.CodeServerBusy, http.StatusInternalServerError)
		}
		return
	}

	httpUtil.ResponseSuccess(c, nil)
}
