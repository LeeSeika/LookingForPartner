package errs

import (
	"errors"
	apierrors "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	DBDuplicatedIdempotencyKey = errors.New("duplicated idempotency key")
)

var (
	RpcUnknown                  = status.Error(codes.Unknown, "unknown")
	RpcNotFound                 = status.Error(codes.NotFound, "not found")
	RpcAlreadyExists            = status.Error(codes.AlreadyExists, "already exists")
	RpcPermissionDenied         = status.Error(codes.PermissionDenied, "permission denied")
	RpcDuplicatedIdempotencyKey = status.Error(codes.AlreadyExists, "duplicated idempotency key")
)

var (
	ApiInternal             = "internal"
	ApiProcessWxLoginFailed = "failed to process wx login"
	ApiNotFound             = "not found"
	ApiUnauthorized         = "unauthorized"
	ApiGenTokenFailed       = "failed to generate token"
	ApiPermissionDenied     = "permission denied"
)

func FormatApiError(statusCode int, errMsg string) error {
	return apierrors.New(statusCode, errMsg)
}

// formatted errs

func FormattedApiInternal() error {
	return apierrors.New(http.StatusInternalServerError, ApiInternal)
}

func FormattedApiNotFound() error {
	return apierrors.New(http.StatusBadRequest, ApiNotFound)
}

func FormattedApiUnAuthorized() error {
	return apierrors.New(http.StatusUnauthorized, ApiUnauthorized)
}

func FormattedApiGenTokenFailed() error {
	return apierrors.New(http.StatusInternalServerError, ApiGenTokenFailed)
}
