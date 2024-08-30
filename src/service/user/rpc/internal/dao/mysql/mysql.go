package mysql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	basedao "lookingforpartner/common/dao"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/user/model"
	"lookingforpartner/service/user/rpc/internal/dao"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) UpdatePostCount(ctx context.Context, wxUid string, delta int, idempotencyKey int64) error {
	tx := m.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// check idempotency
	idempotency := model.IdempotencyUser{
		ID: idempotencyKey,
	}
	rs := tx.Create(idempotency)
	if rs.Error != nil {
		if errors.Is(rs.Error, gorm.ErrDuplicatedKey) {
			return errs.DBDuplicatedIdempotencyKey
		}
		return rs.Error
	}

	rs = tx.Model(&model.User{}).
		Where("wx_uid = ?", wxUid).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", delta))
	if rs.Error != nil {
		return rs.Error
	}

	rs = tx.Commit()

	return rs.Error
}

func (m *MysqlInterface) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	db := m.db.WithContext(ctx)

	query := db.Clauses(clause.Returning{}).
		Where("wx_uid = ?", user.WxUid)

	updateOpt := basedao.UpdateOpt{
		Query: query,
		Data:  user,
	}

	if err := basedao.Update(db, updateOpt); err != nil {
		return nil, err
	}

	return user, nil
}

func (m *MysqlInterface) GetUser(ctx context.Context, wxUid string) (*model.User, error) {
	var user model.User
	rs := m.db.WithContext(ctx).Model(&model.User{}).Where("wx_uid = ?", wxUid).First(&user)
	return &user, rs.Error
}

func (m *MysqlInterface) FirstOrCreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	rs := m.db.WithContext(ctx).Where("wx_uid = ?", user.WxUid).FirstOrCreate(user)
	return user, rs.Error
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
		fmt.Println("failed to open mysql")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to Connect mysql server, err:" + err.Error())
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
	m.db.AutoMigrate(&model.User{})
	m.db.AutoMigrate(&model.IdempotencyUser{})
}
