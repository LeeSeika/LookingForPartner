package svc

import (
	"log"
	"lookingforpartner/service/user/rpc/internal/config"
	"lookingforpartner/service/user/rpc/internal/dao"
	"lookingforpartner/service/user/rpc/internal/dao/mysql"
)

type ServiceContext struct {
	Config        config.Config
	UserInterface dao.UserInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	userInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create user interface, err: %v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:        c,
		UserInterface: userInterface,
	}
}
