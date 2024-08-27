package mysql

import (
	"context"
	"fmt"
	"time"

	basedao "lookingforpartner/common/dao"
	"lookingforpartner/model"
	"lookingforpartner/service/post/rpc/internal/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	var post model.Post
	post.PostID = postID

	// delete post with project
	rs := tx.Select("Project").Delete(&post)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// decrease post_count by 1
	if err := tx.Where("wx_uid = ?", post.AuthorID).
		UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (m *MysqlInterface) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// create post with cascade project
	if err := tx.Create(post).Error; err != nil {
		return nil, err
	}
	if err := tx.Where("id = ?", post.PostID).First(post).Error; err != nil {
		return nil, err
	}

	// increase post_count by 1
	if err := tx.Where("wx_uid = ?", post.AuthorID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (m *MysqlInterface) GetPost(ctx context.Context, postID string) (*model.Post, error) {
	db := m.db.WithContext(ctx)

	var post model.Post

	rs := db.Model(&model.Post{}).
		Preload("Projects").
		Where("post_id = ?", postID).
		First(&post)

	return &post, rs.Error
}

func (m *MysqlInterface) GetPosts(ctx context.Context, page, size int64, order basedao.OrderOpt) ([]*model.Post, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	posts := make([]*model.Post, 0, int(size))

	query := db.Preload("Projects")

	param := basedao.PaginationParam{
		Query:   query,
		Page:    int(page),
		Limit:   int(size),
		OrderBy: []string{order.String()},
		ShowSQL: false,
	}
	paginator, err := basedao.GetListWithPagination(db, &param, posts)

	return posts, paginator, err
}

func (m *MysqlInterface) GetPostsByAuthorID(ctx context.Context, page, size int64, authorID string, order basedao.OrderOpt) ([]*model.Post, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	posts := make([]*model.Post, 0, int(size))

	query := db.Where("author_id = ?", authorID).
		Preload("Projects")

	param := basedao.PaginationParam{
		Query:   query,
		Page:    int(page),
		Limit:   int(size),
		OrderBy: []string{order.String()},
		ShowSQL: false,
	}

	paginator, err := basedao.GetListWithPagination(db, &param, posts)

	return posts, paginator, err
}

func (m *MysqlInterface) UpdateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	db := m.db.WithContext(ctx)

	query := db.Where("project_id = ?", project.ProjectID)

	updateOpt := basedao.UpdateOpt{
		Query: query,
		Data:  project,
	}

	if err := basedao.Update(db, updateOpt); err != nil {
		return nil, err
	}

	// get updated project
	getOpt := basedao.GetOpt{
		Query:   query,
		Preload: nil,
		OrderBy: nil,
	}

	// update progress has succeeded, don't return error
	_ = basedao.GetOne(db, getOpt, project)

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
	m.db.AutoMigrate(&model.Project{})
	m.db.AutoMigrate(&model.Post{})
}
