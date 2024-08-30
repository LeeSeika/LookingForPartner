package mqs

import (
	"context"
	"encoding/json"
	"lookingforpartner/pb/user"
	"lookingforpartner/service/user/rpc/internal/logic"
	"lookingforpartner/service/user/rpc/internal/svc"
)

type UpdateUserPostCount struct {
	svcCtx *svc.ServiceContext
}

func NewUpdateUserPostCount(svcCtx *svc.ServiceContext) *UpdateUserPostCount {
	return &UpdateUserPostCount{
		svcCtx: svcCtx,
	}
}

func (c *UpdateUserPostCount) Consume(ctx context.Context, key, val string) error {
	updateUserPostCountReq := user.UpdateUserPostCountRequest{}
	err := json.Unmarshal([]byte(val), &updateUserPostCountReq)
	if err != nil {
		return err
	}

	l := logic.NewUpdateUserPostCountLogic(ctx, c.svcCtx)

	_, err = l.UpdateUserPostCount(&updateUserPostCountReq)
	if err != nil {
		return err
	}

	return nil
}
