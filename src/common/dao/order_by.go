package dao

import "fmt"

type OrderBy string

const (
	OrderByDESC OrderBy = "desc"
	OrderByASC  OrderBy = "asc"
)

const (
	CreatedAt = "created_at"
	UpdatedAt = "updated_at"
)

// String to string.
func (opt *OrderOpt) String() string {
	if opt.OrderBy == "" {
		opt.OrderBy = OrderByASC
	}
	return fmt.Sprintf("%s %s", opt.ColumnName, opt.OrderBy)
}
