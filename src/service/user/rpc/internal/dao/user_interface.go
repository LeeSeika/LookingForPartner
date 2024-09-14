package dao

import (
	"context"
	"lookingforpartner/service/user/model"
)

type UserInterface interface {
	FirstOrCreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, wxUid string) (*model.User, error)
	UpdatePostCount(ctx context.Context, wxUid string, delta int, idempotencyKey int64) error
}
