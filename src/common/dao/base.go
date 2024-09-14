package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"reflect"
)

// InsertOpt insert table option.
type InsertOpt struct {
	// Use gorm table struct.
	Data interface{}
}

// GetOpt query table option.
type GetOpt struct {
	// Use gorm table struct.
	Query   interface{}
	Preload []string
	OrderBy []OrderOpt
}

// UpdateOpt update table option.
type UpdateOpt struct {
	// Use gorm table struct.
	Query interface{}

	// Update attributes with `struct`, will only update non-zero value fields.
	// with `map` to update zero value fields.
	Data interface{}
}

// DeleteOpt delete table option.
type DeleteOpt struct {
	// Use gorm table struct.
	Query interface{}
}

// CountOpt count table option.
type CountOpt struct {
	// Use gorm table struct.
	Query interface{}
}

// OrderOpt query db order by option.
type OrderOpt struct {
	ColumnName string
	OrderBy    OrderBy
}

// Insert data to table.
func Insert(db *gorm.DB, opt InsertOpt) error {
	if err := db.Model(opt.Data).Create(opt.Data).Error; err != nil {
		return errors.Wrap(err, "insert")
	}

	return nil
}

// GetOne data by option.
func GetOne(db *gorm.DB, opt GetOpt, dest interface{}) error {
	query := db.Where(opt.Query)

	if query.Error != nil {
		return errors.Wrapf(query.Error, "db where:%+v", opt.Query)
	}

	for _, o := range opt.OrderBy {
		query = query.Order(o.String())
		if query.Error != nil {
			return errors.Wrapf(query.Error, "db order by:%v", o)
		}
	}

	for _, p := range opt.Preload {
		query = query.Preload(p)
		if query.Error != nil {
			return errors.Wrapf(query.Error, "db preload:%v", p)
		}
	}

	if err := query.First(dest).Error; err != nil {
		return err
	}

	return nil
}

// GetList gets list by option.
func GetList(db *gorm.DB, opt GetOpt, dest interface{}) error {
	query := db.Where(opt.Query)

	if query.Error != nil {
		return errors.Wrapf(query.Error, "db where:%+v", opt.Query)
	}

	for _, o := range opt.OrderBy {
		query = query.Order(o.String())
		if query.Error != nil {
			return errors.Wrapf(query.Error, "db order by:%v", o)
		}
	}

	for _, p := range opt.Preload {
		query = query.Preload(p)
		if query.Error != nil {
			return errors.Wrapf(query.Error, "db preload:%v", p)
		}
	}

	if err := query.Find(dest).Error; err != nil {
		return errors.Wrap(err, "query")
	}

	rvPtr := reflect.ValueOf(dest)
	if rvPtr.Kind() == reflect.Ptr {
		rv := rvPtr.Elem()
		if rv.Kind() == reflect.Slice {
			len := rv.Len()
			if len == 0 {
				return gorm.ErrRecordNotFound
			}
		}
	}

	return nil
}

// Update by option.
func Update(db *gorm.DB, opt UpdateOpt) error {
	query := db.Where(opt.Query)

	if query.Error != nil {
		return errors.Wrapf(query.Error, "db where:%+v", opt.Query)
	}

	query = query.Updates(opt.Data)
	if query.Error != nil {
		return errors.Wrap(query.Error, "update")
	}

	return nil
}

// Delete by option.
func Delete(db *gorm.DB, opt DeleteOpt) error {
	query := db.Where(opt.Query)

	if query.Error != nil {
		return errors.Wrapf(query.Error, "db where:%+v", opt.Query)
	}

	query = query.Delete(opt.Query)
	if query.Error != nil {
		return errors.Wrap(query.Error, "delete")
	}

	return nil
}

// Count by option.
func Count(db *gorm.DB, opt CountOpt) (int64, error) {
	var count int64

	query := db.Where(opt.Query)

	if query.Error != nil {
		return count, errors.Wrapf(query.Error, "db where:%+v", opt.Query)
	}

	if err := query.Model(opt.Query).Count(&count).Error; err != nil {
		return count, errors.Wrap(err, "count")
	}

	return count, nil
}
