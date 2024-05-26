package errs

import (
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	RpcUnknown  = status.Error(codes.Unknown, "unknown")
	RpcNotFound = status.Error(codes.NotFound, "not found")
)

var (
	ApiUnknown              = "unknown"
	ApiProcessWxLoginFailed = "failed to process wx login"
	ApiNotFound             = "not found"
	ApiUnauthorized         = "unauthorized"
	ApiGenTokenFailed       = "failed to generate token"
)

func FormatApiError(statusCode int, errMsg string) error {
	return errors.New(statusCode, errMsg)
}

// formatted errs

func FormattedUnknown() error {
	return errors.New(http.StatusInternalServerError, ApiUnknown)
}

func FormattedNotFound() error {
	return errors.New(http.StatusBadRequest, ApiNotFound)
}

func FormattedUnAuthorized() error {
	return errors.New(http.StatusUnauthorized, ApiUnauthorized)
}

func FormattedGenTokenFailed() error {
	return errors.New(http.StatusInternalServerError, ApiGenTokenFailed)
}
