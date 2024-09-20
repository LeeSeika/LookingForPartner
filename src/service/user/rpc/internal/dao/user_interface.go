package dao

import (
	"context"
	"lookingforpartner/service/user/model/entity"
)

type UserInterface interface {
	FirstOrCreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, wxUid string) (*entity.User, error)
	UpdatePostCount(ctx context.Context, wxUid string, delta int, idempotencyKey int64) error
}
