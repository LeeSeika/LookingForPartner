package mysql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm/clause"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/model"
	"time"

	basedao "lookingforpartner/common/dao"
	"lookingforpartner/service/post/rpc/internal/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) DeletePost(ctx context.Context, postID string, idempotencyKey int64) (*model.PostProject, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// check idempotency
	idempotency := model.IdempotencyPost{
		ID: idempotencyKey,
	}
	rs := tx.Create(idempotency)
	if rs.Error != nil {
		if errors.Is(rs.Error, gorm.ErrDuplicatedKey) {
			return nil, errs.DBDuplicatedIdempotencyKey
		}
		return nil, rs.Error
	}

	// delete post
	var post model.Post
	post.PostID = postID
	rs = tx.Clauses(clause.Returning{}).Delete(&post)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// delete project associated with post
	var proj model.Project
	proj.PostID = postID
	rs = tx.Clauses(clause.Returning{}).Delete(&proj)
	if rs.Error != nil {
		return nil, rs.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	poProj := model.PostProject{Post: &post, Project: &proj}

	return &poProj, nil
}

func (m *MysqlInterface) CreatePost(ctx context.Context, post *model.Post, proj *model.Project, idempotencyKey int64) (*model.PostProject, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// check idempotency
	idempotency := model.IdempotencyPost{
		ID: idempotencyKey,
	}
	rs := tx.Create(idempotency)
	if rs.Error != nil {
		if errors.Is(rs.Error, gorm.ErrDuplicatedKey) {
			return nil, errs.DBDuplicatedIdempotencyKey
		}
		return nil, rs.Error
	}

	// create post
	if err := tx.Create(post).Error; err != nil {
		return nil, err
	}
	if err := tx.Create(proj).Error; err != nil {
		return nil, err
	}

	if err := tx.First(post).Error; err != nil {
		return nil, err
	}
	if err := tx.First(proj).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	poProj := model.PostProject{Post: post, Project: proj}

	return &poProj, nil
}

func (m *MysqlInterface) GetPost(ctx context.Context, postID string) (*model.PostProject, error) {
	db := m.db.WithContext(ctx)

	var poProj model.PostProject

	rs := db.Model(&model.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id").
		Where("posts.post_id = ?", postID).
		First(&poProj)

	return &poProj, rs.Error
}

func (m *MysqlInterface) GetPosts(ctx context.Context, page, size int64, order basedao.OrderOpt) ([]*model.PostProject, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	poProjs := make([]*model.PostProject, 0, int(size))

	query := db.Model(&model.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id")

	param := basedao.PaginationParam{
		Query:   query,
		Page:    int(page),
		Limit:   int(size),
		OrderBy: []string{order.String()},
		ShowSQL: false,
	}
	paginator, err := basedao.GetListWithPagination(db, &param, poProjs)

	return poProjs, paginator, err
}

func (m *MysqlInterface) GetPostsByAuthorID(ctx context.Context, page, size int64, authorID string, order basedao.OrderOpt) ([]*model.PostProject, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	poProjs := make([]*model.PostProject, 0, int(size))

	query := db.Model(&model.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id").
		Where("author_id = ?", authorID)

	param := basedao.PaginationParam{
		Query:   query,
		Page:    int(page),
		Limit:   int(size),
		OrderBy: []string{order.String()},
		ShowSQL: false,
	}

	paginator, err := basedao.GetListWithPagination(db, &param, poProjs)

	return poProjs, paginator, err
}

func (m *MysqlInterface) UpdateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	db := m.db.WithContext(ctx)

	query := db.Clauses(clause.Returning{}).
		Where("project_id = ?", project.ProjectID)

	updateOpt := basedao.UpdateOpt{
		Query: query,
		Data:  project,
	}

	if err := basedao.Update(db, updateOpt); err != nil {
		return nil, err
	}

	return project, nil
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (dao.PostInterface, error) {
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
	m.db.AutoMigrate(&model.Project{})
	m.db.AutoMigrate(&model.Post{})
	m.db.AutoMigrate(&model.IdempotencyPost{})
}
