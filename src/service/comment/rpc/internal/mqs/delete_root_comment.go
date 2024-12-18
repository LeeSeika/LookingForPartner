package mqs

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"lookingforpartner/common/event"
	"lookingforpartner/pb/comment"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/logic"
	"lookingforpartner/service/comment/rpc/internal/svc"
)

type DeleteRootComment struct {
	svcCtx *svc.ServiceContext
	conf   config.Config
	logger logx.Logger
}

func NewDeleteRootComment(conf config.Config, svcCtx *svc.ServiceContext) *DeleteRootComment {
	return &DeleteRootComment{
		svcCtx: svcCtx,
		conf:   conf,
		logger: logx.WithContext(context.Background()).WithFields(logx.Field("consumer-topic", conf.KqDeleteRootCommentConsumerConf.Topic)),
	}
}

func (c *DeleteRootComment) Consume(ctx context.Context, key, val string) error {
	var evt event.DeleteRootComment
	if err := json.Unmarshal([]byte(val), &evt); err != nil {
		c.logger.
			Errorf("cannot unmarshal json, err: %+v", err)
		return err
	}

	deleteSubCommentsByRootIDReq := comment.DeleteSubCommentsByRootIDRequest{RootID: evt.CommentID}
	l := logic.NewDeleteSubCommentsByRooIDLogic(ctx, c.svcCtx)
	_, err := l.DeleteSubCommentsByRooID(&deleteSubCommentsByRootIDReq)
	if err != nil {
		c.logger.Errorf("cannot delete sub comments by root id, err: %+v", err)
		return err
	}

	return nil
}
