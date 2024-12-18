package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

func SetupLogger(service string) {
	path := "/var/logs/" + service
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}
	conf := logx.LogConf{
		ServiceName: service,
		Mode:        "file",
		Path:        path,
		KeepDays:    7,
		MaxBackups:  2,
		Rotation:    "daily",
		Stat:        true,
	}
	logx.MustSetup(conf)
}
