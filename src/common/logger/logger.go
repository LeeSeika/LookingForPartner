package logger

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewLogger(ctx context.Context, service string) logx.Logger {
	conf := logx.LogConf{
		ServiceName: service,
		Mode:        "file",
		Path:        "logs/" + service,
		KeepDays:    7,
		MaxBackups:  2,
		Rotation:    "daily",
	}
	logx.MustSetup(conf)
	return logx.WithContext(ctx)
}
