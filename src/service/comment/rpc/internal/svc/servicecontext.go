package svc

import (
	"log"
	"lookingforpartner/service/comment/rpc/internal/config"
	"lookingforpartner/service/comment/rpc/internal/dao"
	"lookingforpartner/service/comment/rpc/internal/dao/mysql"
)

type ServiceContext struct {
	Config           config.Config
	CommentInterface dao.CommentInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	commentInterface, err := mysql.NewMysqlInterface(c.Mysql.DataBase, c.Mysql.Username, c.Mysql.Password, c.Mysql.Host, c.Mysql.Port, c.Mysql.MaxIdleConns, c.Mysql.MaxOpenConns, c.Mysql.ConnMaxLifeTime)
	if err != nil {
		log.Printf("failed to create post interface, err: %v\n", err)
		return nil
	}
	return &ServiceContext{
		Config:           c,
		CommentInterface: commentInterface,
	}
}