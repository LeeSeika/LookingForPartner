package logic

import (
	"context"

	"lookingforpartner/pb/leaf"
	"lookingforpartner/service/leaf/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type NextSegmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNextSegmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NextSegmentLogic {
	return &NextSegmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NextSegmentLogic) NextSegment(in *leaf.NextSegmentRequest) (*leaf.NextSegmentResponse, error) {
	// todo: add your logic here and delete this line

	return &leaf.NextSegmentResponse{}, nil
}
