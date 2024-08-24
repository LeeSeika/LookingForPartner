package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model2 "lookingforpartner/model"
	"lookingforpartner/service/user/rpc/internal/dao"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) SetUser(user *model2.User) error {
	rs := m.db.Where("wx_uid = ?", user.WxUid).Updates(user)
	return rs.Error
}

func (m *MysqlInterface) GetUser(wxUid string) (*model2.User, error) {
	var user model2.User
	rs := m.db.Model(&model2.User{}).Where("wx_uid = ?", wxUid).First(&user)
	return &user, rs.Error
}

func (m *MysqlInterface) FirstOrCreateUser(user *model2.User) error {
	rs := m.db.Where("wx_uid = ?", user.WxUid).FirstOrCreate(user)
	return rs.Error
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (dao.UserInterface, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open mysql")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to Connect mysql server, err:" + err.Error())
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(connMaxLifeTime))

	m := &MysqlInterface{
		db: db,
	}
	m.autoMigrate()
	return m, nil
}

func (m *MysqlInterface) autoMigrate() {
	m.db.AutoMigrate(&model2.User{})
}
