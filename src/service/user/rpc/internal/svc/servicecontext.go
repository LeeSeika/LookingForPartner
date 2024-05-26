package svc

import (
	"log"
	"lookingforpartner/service/user/model"
	"lookingforpartner/service/user/model/mysql"
	"lookingforpartner/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserInterface model.UserInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	userInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("Failed to create user interface, err: %v\n", err)
		return nil
	}

	return &ServiceContext{
		Config:        c,
		UserInterface: userInterface,
	}
}
