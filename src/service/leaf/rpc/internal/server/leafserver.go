// Code generated by goctl. DO NOT EDIT.
// Source: leaf.proto

package server

import (
	"context"

	"lookingforpartner/pb/leaf"
	"lookingforpartner/service/leaf/rpc/internal/logic"
	"lookingforpartner/service/leaf/rpc/internal/svc"
)

type LeafServer struct {
	svcCtx *svc.ServiceContext
	leaf.UnimplementedLeafServer
}

func NewLeafServer(svcCtx *svc.ServiceContext) *LeafServer {
	return &LeafServer{
		svcCtx: svcCtx,
	}
}

func (s *LeafServer) NextSegment(ctx context.Context, in *leaf.NextSegmentRequest) (*leaf.NextSegmentResponse, error) {
	l := logic.NewNextSegmentLogic(ctx, s.svcCtx)
	return l.NextSegment(in)
}
