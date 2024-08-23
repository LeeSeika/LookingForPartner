package dao

import (
	"math"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// PaginationParam parameters for pagination
type PaginationParam struct {
	Query   interface{}
	Page    int
	Limit   int
	OrderBy []string
	ShowSQL bool
}

// Paginator indicates the result of paging
type Paginator struct {
	TotalRecord int64 `json:"total_record"`
	TotalPage   int   `json:"total_page"`
	Offset      int   `json:"offset"`
	Limit       int   `json:"limit"`
	CurrPage    int   `json:"curr_page"`
	PrevPage    int   `json:"prev_page"`
	NextPage    int   `json:"next_page"`
}

func GetListWithPagination(db *gorm.DB, p *PaginationParam, result interface{}) (*Paginator, error) {
	query := db.Where(p.Query)
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
	var paginator Paginator
	var count int64
	var offset int

	go countRecords(query, result, done, &count)

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

	paginator.TotalRecord = count
	paginator.CurrPage = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	return &paginator, nil
}

func countRecords(db *gorm.DB, anyType interface{}, done chan error, count *int64) {
	rs := db.Model(anyType).Count(count)
	if rs.Error != nil {
		done <- rs.Error
		return
	}

	close(done)
}
