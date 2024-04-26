package svc

import (
	"log"
	"lookingforpartner/common/snowflake"
	"lookingforpartner/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	err := snowflake.Init(c.Snowflake.MachineID)
	if err != nil {
		log.Printf("failed to initialize snowflake, err:%v\n", err)
		return nil
	}

	return &ServiceContext{
		Config: c,
	}
}
