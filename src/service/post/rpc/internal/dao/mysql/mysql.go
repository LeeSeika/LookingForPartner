package mysql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm/clause"
	"lookingforpartner/common/errs"
	"lookingforpartner/service/post/model/entity"
	"lookingforpartner/service/post/model/vo"
	"time"

	basedao "lookingforpartner/common/dao"
	"lookingforpartner/service/post/rpc/internal/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) UpdatePost(ctx context.Context, updatedPost *entity.Post) (*entity.Post, error) {
	db := m.db.WithContext(ctx)

	if err := db.Save(updatedPost).Error; err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (m *MysqlInterface) DeletePost(ctx context.Context, postID string) (*vo.PostProject, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// delete post
	var post entity.Post
	post.PostID = postID
	rs := tx.Clauses(clause.Returning{}).Delete(&post)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// delete project associated with post
	var proj entity.Project
	proj.PostID = postID
	rs = tx.Clauses(clause.Returning{}).Delete(&proj)
	if rs.Error != nil {
		return nil, rs.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	poProj := vo.PostProject{Post: &post, Project: &proj}

	return &poProj, nil
}

func (m *MysqlInterface) CreatePost(ctx context.Context, post *entity.Post, proj *entity.Project, idempotencyKey int64) (*vo.PostProject, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// check idempotency
	idempotency := entity.IdempotencyPost{
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

	post = &entity.Post{
		PostID: post.PostID,
	}
	proj = &entity.Project{
		ProjectID: proj.ProjectID,
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

	poProj := vo.PostProject{Post: post, Project: proj}

	return &poProj, nil
}

func (m *MysqlInterface) GetPost(ctx context.Context, postID string) (*vo.PostProject, error) {
	db := m.db.WithContext(ctx)

	var poProj vo.PostProject

	rs := db.Model(&entity.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id").
		Where("posts.post_id = ?", postID).
		First(&poProj)

	return &poProj, rs.Error
}

func (m *MysqlInterface) GetPosts(ctx context.Context, page, size int64, order basedao.OrderOpt) ([]*vo.PostProject, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	poProjs := make([]*vo.PostProject, 0, int(size))

	query := db.Model(&entity.Post{}).
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

func (m *MysqlInterface) GetPostsByAuthorID(ctx context.Context, page, size int64, authorID string, order basedao.OrderOpt) ([]*vo.PostProject, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	poProjs := make([]*vo.PostProject, 0, int(size))

	query := db.Model(&entity.Post{}).
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

func (m *MysqlInterface) UpdateProject(ctx context.Context, project *entity.Project) (*entity.Project, error) {
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
	m.db.AutoMigrate(
		&entity.Project{},
		&entity.Post{},
		&entity.IdempotencyPost{},
	)
}
