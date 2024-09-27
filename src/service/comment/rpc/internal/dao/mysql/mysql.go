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
	"lookingforpartner/service/comment/model/entity"
	"lookingforpartner/service/comment/model/vo"
	"lookingforpartner/service/comment/rpc/internal/dao"
	"time"
)

type MysqlInterface struct {
	db *gorm.DB
}

func (m *MysqlInterface) DeleteAllCommentsBySubjectID(ctx context.Context, subjectID string) error {
	db := m.db.WithContext(ctx)

	if err := db.Where("subject_id = ?", subjectID).
		Delete(&entity.CommentIndex{}).Error; err != nil {
		return err
	}

	return nil
}

func (m *MysqlInterface) DeleteSubCommentsByRootID(ctx context.Context, rootID string) error {
	db := m.db.WithContext(ctx)

	if err := db.Where("root_id = ?", rootID).
		Delete(&entity.CommentIndex{}).Error; err != nil {
		return err
	}

	return nil
}

func (m *MysqlInterface) GetSubject(ctx context.Context, subjectID string) (*entity.Subject, error) {
	db := m.db.WithContext(ctx)

	subject := entity.Subject{SubjectID: subjectID}

	if err := db.Model(&entity.Subject{}).First(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func (m *MysqlInterface) GetTopSubCommentsByRootIDs(ctx context.Context, rootIDs []string, topCount int, order basedao.OrderOpt) ([]*vo.CommentIndexContent, error) {
	db := m.db.WithContext(ctx)

	// query sub comment indexes
	subCommentIndexes := make([]entity.CommentIndex, 0, len(rootIDs)*topCount)

	if err := db.Model(&entity.CommentIndex{}).
		Where("root_id IN ?", rootIDs).
		Group("root_id").
		Order(order).
		Limit(topCount).
		Find(&subCommentIndexes).Error; err != nil {
		return nil, err
	}

	// query sub comment contents
	subCommentIDs := make([]string, 0, len(subCommentIndexes))
	for _, subComment := range subCommentIndexes {
		subCommentIDs = append(subCommentIDs, subComment.CommentID)
	}

	subCommentContents := make([]entity.CommentContent, 0, len(subCommentIndexes))

	if err := db.Model(&entity.CommentContent{}).
		Where("comment_id IN ?", subCommentIDs).
		Find(&subCommentContents).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: "FIELD(comment_id,?)", Vars: []interface{}{subCommentIDs}, WithoutParentheses: true},
		}).Error; err != nil {
		return nil, err
	}

	// construct vo
	subComments := make([]*vo.CommentIndexContent, 0, len(subCommentIndexes))
	for i := 0; i < len(subCommentIndexes); i++ {
		subComment := vo.CommentIndexContent{
			CommentIndex:   &subCommentIndexes[i],
			CommentContent: &subCommentContents[i],
		}
		subComments = append(subComments, &subComment)
	}

	return subComments, nil
}

func (m *MysqlInterface) CreateComment(ctx context.Context, commentIndex *entity.CommentIndex, commentContent *entity.CommentContent) (*vo.CommentIndexContent, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// assign a floor number for comment
	if commentIndex.RootID != nil {
		// if this is a sub comment, its floor is 0
		commentIndex.Floor = 0
	} else {
		// if this is a root comment, assign a floor
		var subject entity.Subject
		subject.SubjectID = commentIndex.SubjectID
		if err := tx.First(&subject).Error; err != nil {
			return nil, err
		}

		var nextFloor int
		if subject.RootCommentCount == 0 {
			// this is the first comment
			nextFloor = 2
		} else {
			// this is not the first comment, next floor = max floor + 1
			var maxFloor int
			if err := tx.Model(&entity.CommentIndex{}).
				Select("max(floor) AS max_floor").
				Where("subject_id = ?", subject.ID).
				Find(&maxFloor).Error; err != nil {
				return nil, err
			}

			nextFloor = maxFloor + 1
		}
		commentIndex.Floor = nextFloor
	}

	// create comment
	if err := tx.Create(commentIndex).Error; err != nil {
		return nil, err
	}
	if err := tx.Create(commentContent).Error; err != nil {
		return nil, err
	}

	// update subject
	// update all comment count of subject
	query := tx.Where("subject_id = ?", commentIndex.SubjectID).
		UpdateColumn("all_comment_count", gorm.Expr("all_comment_count + ?", 1))
	if commentIndex.RootID == nil {
		// if this is a root comment, also update root comment count of subject
		query = query.UpdateColumn("root_comment_count", gorm.Expr("root_comment_count + ?", 1))
	}
	updateOpt := basedao.UpdateOpt{
		Query: query,
		Data:  entity.Subject{},
	}
	if err := basedao.Update(tx, updateOpt); err != nil {
		return nil, err
	}

	// if this is a sub comment, update sub comments count of root comment
	if commentIndex.RootID != nil {
		query := tx.Where("comment_id = ?", commentIndex.RootID).
			UpdateColumn("sub_comment_count", gorm.Expr("sub_comment_count + ?", 1))
		updateOpt := basedao.UpdateOpt{
			Query: query,
			Data:  entity.CommentIndex{},
		}
		if err := basedao.Update(tx, updateOpt); err != nil {
			return nil, err
		}
	}

	commentIndex = &entity.CommentIndex{CommentID: commentIndex.CommentID}
	commentContent = &entity.CommentContent{CommentID: commentContent.CommentID}

	if err := tx.First(commentIndex).Error; err != nil {
		return nil, err
	}
	if err := tx.First(commentContent).Error; err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	commentIndexContent := vo.CommentIndexContent{
		CommentIndex:   commentIndex,
		CommentContent: commentContent,
	}

	return &commentIndexContent, nil
}

func (m *MysqlInterface) GetComment(ctx context.Context, commentID string) (*vo.CommentIndexContent, error) {
	db := m.db.WithContext(ctx)

	var commentIndexContent vo.CommentIndexContent
	rs := db.Model(&entity.CommentIndex{}).
		Joins("left join comment_contents on comment_indexes.comment_id = comment_contents.comment_id").
		Where("comment_id = ?", commentID).
		First(&commentIndexContent)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return &commentIndexContent, nil
}

func (m *MysqlInterface) GetRootCommentsByPostID(ctx context.Context, postID string, page, size int64, order basedao.OrderOpt) ([]*vo.CommentIndexContent, *basedao.Paginator, error) {
	db := m.db.WithContext(ctx)

	rootCommentIndexContents := make([]*vo.CommentIndexContent, 0, int(size))

	// query root comments
	queryRootComments := db.Model(&entity.CommentIndex{}).
		Joins("left join comment_contents on comment_indexes.comment_id = comment_contents.comment_id").
		Where("post_id == ?", postID).
		Where("comment_indexes.root_id == NULL")

	pagiParam := basedao.PaginationParam{
		Query:   queryRootComments,
		Page:    int(page),
		Limit:   int(size),
		OrderBy: []string{order.String()},
		ShowSQL: false,
	}

	paginator, err := basedao.GetListWithPagination(db, &pagiParam, &rootCommentIndexContents)
	if err != nil {
		return nil, nil, err
	}

	return rootCommentIndexContents, paginator, nil
}

func (m *MysqlInterface) DeleteComment(ctx context.Context, commentID string) (*vo.CommentIndexContent, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// delete comment index
	var commentIndex entity.CommentIndex
	commentIndex.CommentID = commentID
	rs := tx.Clauses(clause.Returning{}).Delete(&commentIndex)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// delete comment content
	var commentContent entity.CommentContent
	commentContent.CommentID = commentID
	rs = tx.Clauses(clause.Returning{}).Delete(&commentContent)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// if this is a sub comment, update sub comments count of its root comment
	if commentIndex.RootID != nil {
		query := tx.Where("comment_id = ?", commentIndex.RootID).
			UpdateColumn("sub_comment_count", gorm.Expr("sub_comment_count - ?", 1))
		updateOpt := basedao.UpdateOpt{
			Query: query,
			Data:  entity.CommentIndex{},
		}
		if err := basedao.Update(tx, updateOpt); err != nil {
			return nil, err
		}
	}

	// update subject
	query := tx.Where("subject_id = ?", commentIndex.SubjectID)
	var allCommentCountGoingToBeDeleted int64 = 1
	if commentIndex.RootID == nil {
		// if this is a root comment, update root comment count of subject and all comment count
		query = query.UpdateColumn("root_comment_count", gorm.Expr("root_comment_count - ?", 1))
		// query sub comment count
		tx.Model(&entity.CommentIndex{}).
			Where("root_id = ?", commentIndex.CommentID).
			Count(&allCommentCountGoingToBeDeleted)
	}
	// update all comment count of subject
	query = query.
		UpdateColumn("all_comment_count", gorm.Expr("all_comment_count - ?", allCommentCountGoingToBeDeleted))
	updateOpt := basedao.UpdateOpt{
		Query: query,
		Data:  entity.Subject{},
	}
	if err := basedao.Update(tx, updateOpt); err != nil {
		return nil, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	commentIndexContent := vo.CommentIndexContent{CommentIndex: &commentIndex, CommentContent: &commentContent}
	return &commentIndexContent, nil
}

func (m *MysqlInterface) CreateSubject(ctx context.Context, subject *entity.Subject, idempotencyKey int64) (*entity.Subject, error) {
	tx := m.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer tx.Rollback()

	tx = tx.WithContext(ctx)

	// check idempotency
	idempotency := entity.IdempotencyComment{
		ID: idempotencyKey,
	}
	rs := tx.Create(idempotency)
	if rs.Error != nil {
		if errors.Is(rs.Error, gorm.ErrDuplicatedKey) {
			return nil, errs.DBDuplicatedIdempotencyKey
		}
		return nil, rs.Error
	}

	// create subject
	if err := tx.Model(&entity.Subject{}).Create(subject).Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&entity.Subject{}).
		Where("subject_id = ?", subject.SubjectID).
		First(subject).Error; err != nil {
		// return original subject, without error
		return subject, nil
	}

	return subject, nil
}

func (m *MysqlInterface) UpdateSubject(ctx context.Context, updatedSubject *entity.Subject) (*entity.Subject, error) {
	db := m.db.WithContext(ctx)

	if err := db.Save(updatedSubject).Error; err != nil {
		return nil, err
	}

	return updatedSubject, nil
}

func (m *MysqlInterface) DeleteSubject(ctx context.Context, subjectID string) (*entity.Subject, error) {
	db := m.db.WithContext(ctx)

	subject := entity.Subject{SubjectID: subjectID}
	if err := db.Clauses(clause.Returning{}).Delete(&subject).Error; err != nil {
		return nil, err
	}

	return &subject, nil
}

func NewMysqlInterface(database, username, password, host, port string, maxIdleConns, maxOpenConns, connMaxLifeTime int) (dao.CommentInterface, error) {
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
		&entity.CommentIndex{},
		&entity.CommentContent{},
		&entity.Subject{},
		&entity.IdempotencyComment{},
	)
}
