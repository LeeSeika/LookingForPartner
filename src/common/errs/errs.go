package errs

import (
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	RpcUnknown          = status.Error(codes.Unknown, "unknown")
	RpcNotFound         = status.Error(codes.NotFound, "not found")
	RpcAlreadyExists    = status.Error(codes.AlreadyExists, "already exists")
	RpcPermissionDenied = status.Error(codes.PermissionDenied, "permission denied")
)

var (
	ApiInternal             = "internal"
	ApiProcessWxLoginFailed = "failed to process wx login"
	ApiNotFound             = "not found"
	ApiUnauthorized         = "unauthorized"
	ApiGenTokenFailed       = "failed to generate token"
)

func FormatApiError(statusCode int, errMsg string) error {
	return errors.New(statusCode, errMsg)
}

// formatted errs

func FormattedApiInternal() error {
	return errors.New(http.StatusInternalServerError, ApiInternal)
}

func FormattedApiNotFound() error {
	return errors.New(http.StatusBadRequest, ApiNotFound)
}

func FormattedApiUnAuthorized() error {
	return errors.New(http.StatusUnauthorized, ApiUnauthorized)
}

func FormattedApiGenTokenFailed() error {
	return errors.New(http.StatusInternalServerError, ApiGenTokenFailed)
}
