package dao

import (
	"lookingforpartner/pb/paginator"
	"math"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// PaginationParam parameters for pagination
type PaginationParam struct {
	DB      *gorm.DB
	Query   *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

// Paginator indicates the result of paging
type Paginator struct {
	TotalRecord int64
	TotalPage   int
	Offset      int
	Limit       int
	CurrPage    int
	PrevPage    int
	NextPage    int
}

func GetListWithPagination(p *PaginationParam, result interface{}) (*Paginator, error) {
	query := p.Query
	if query.Error != nil {
		return nil, errors.Wrapf(query.Error, "db where:%+v", query)

	}

	if p.ShowSQL {
		query = query.Debug()
	}
	if p.Page < 1 {
		p.Page = DefaultPageNumber
	}
	if p.Limit == 0 {
		p.Limit = DefaultSizeNumber
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			query = query.Order(o)
			if query.Error != nil {
				return nil, errors.Wrapf(query.Error, "db order by:%v", o)
			}
		}
	}

	done := make(chan error, 1)
	var pagi Paginator
	var count int64
	var offset int

	go countRecords(p.DB, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	query = query.Limit(p.Limit).Offset(offset).Find(result)
	if query.Error != nil {
		return nil, errors.Wrapf(query.Error, "db find:%+v", query)
	}

	// check the count goroutine
	err, valid := <-done
	if valid {
		return nil, errors.Wrapf(err, "db failed to count")
	}

	pagi.TotalRecord = count
	pagi.CurrPage = p.Page

	pagi.Offset = offset
	pagi.Limit = p.Limit
	pagi.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		pagi.PrevPage = p.Page - 1
	} else {
		pagi.PrevPage = p.Page
	}

	if p.Page == pagi.TotalPage {
		pagi.NextPage = p.Page
	} else {
		pagi.NextPage = p.Page + 1
	}

	return &pagi, nil
}

func countRecords(db *gorm.DB, anyType interface{}, done chan error, count *int64) {
	rs := db.Model(anyType).Count(count)
	if rs.Error != nil {
		done <- rs.Error
		return
	}

	close(done)
}

func (p *Paginator) ToRPC() *paginator.Paginator {
	return &paginator.Paginator{
		TotalRecord: p.TotalRecord,
		TotalPage:   int64(p.TotalPage),
		Offset:      int64(p.Offset),
		Limit:       int64(p.Limit),
		CurrPage:    int64(p.CurrPage),
		PrevPage:    int64(p.PrevPage),
		NextPage:    int64(p.NextPage),
	}
}
