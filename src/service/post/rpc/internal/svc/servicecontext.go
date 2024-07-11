package svc

import (
	"log"
	"lookingforpartner/service/post/model"
	"lookingforpartner/service/post/model/mysql"
	"lookingforpartner/service/post/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	PostInterface model.PostInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	postInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("Failed to create post interface, err: %v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:        c,
		PostInterface: postInterface,
	}
}
