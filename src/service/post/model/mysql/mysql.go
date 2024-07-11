package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lookingforpartner/common/dao"
	"lookingforpartner/service/post/model"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) DeletePostTx(postID int64) (*model.Post, *model.Project, error) {
	tx := m.db.Begin()
	defer tx.Rollback()

	var post model.Post
	rs := tx.Where("post_id = ?", postID).Delete(&post)
	if rs.Error != nil {
		return nil, nil, rs.Error
	}

	var project model.Project
	rs = tx.Where("post_id = ?", postID).Delete(&project)
	if rs.Error != nil {
		return nil, nil, rs.Error
	}

	if err := tx.Where("wx_uid = ?", post.AuthorID).
		UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error; err != nil {
		return nil, nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}

	return &post, &project, nil
}

func (m *MysqlInterface) CreatePostWithProjectTx(post *model.Post, project *model.Project) (*model.Post, *model.Project, error) {
	tx := m.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(post).Error; err != nil {
		return nil, nil, err
	}
	if err := tx.Create(project).Error; err != nil {
		return nil, nil, err
	}
	if err := tx.Where("id = ?", post.ID).First(post).Error; err != nil {
		return nil, nil, err
	}
	if err := tx.Where("id = ?", project.ID).First(project).Error; err != nil {
		return nil, nil, err
	}
	if err := tx.Where("wx_uid = ?", post.AuthorID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error; err != nil {
		return nil, nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, nil, err
	}
	return post, project, nil
}

func (m *MysqlInterface) CreatePost(post *model.Post) (*model.Post, error) {
	rs := m.db.Create(post)
	if rs.Error != nil {
		return nil, rs.Error
	}
	_ = m.db.Model(&model.Post{}).Where("id = ?", post.ID).First(post)
	return post, nil

}

func (m *MysqlInterface) GetPost(postID int64) (*model.PostWithProject, error) {
	var postWithProject model.PostWithProject
	rs := m.db.Model(&model.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id").
		Where("posts.post_id = ?", postID).
		First(&postWithProject)
	return &postWithProject, rs.Error
}

func (m *MysqlInterface) GetPosts(page, size int64, order dao.OrderOpt) ([]*model.PostWithProject, error) {
	offset := (page - 1) * size
	limit := size
	postWithProjects := make([]*model.PostWithProject, 0)
	rs := m.db.Model(&model.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id").
		Offset(int(offset)).
		Limit(int(limit)).
		Order(order.String()).
		Find(&postWithProjects)
	return postWithProjects, rs.Error
}

func (m *MysqlInterface) GetPostsByAuthorID(page, size int64, authorID string, order dao.OrderOpt) ([]*model.PostWithProject, error) {
	offset := (page - 1) * size
	limit := size
	postWithProjects := make([]*model.PostWithProject, 0)
	rs := m.db.Model(&model.Post{}).
		Joins("left join projects on posts.post_id = projects.post_id").
		Where("posts.author_id = ?", authorID).
		Offset(int(offset)).
		Limit(int(limit)).
		Order(order.String()).
		Find(&postWithProjects)
	return postWithProjects, rs.Error
}

func (m *MysqlInterface) SetProject(project *model.Project) (*model.Project, error) {
	if err := m.db.Model(&model.Project{}).Where("project_id = ?", project.ProjectID).Updates(project).Error; err != nil {
		return nil, err
	}
	m.db.Model(&model.Project{}).Where("project_id = ?", project.ProjectID).First(project)
	return project, nil
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (model.PostInterface, error) {
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
	m.db.AutoMigrate(&model.Post{})
	m.db.AutoMigrate(&model.Project{})
}
