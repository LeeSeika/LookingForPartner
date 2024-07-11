package dao

import (
	"fmt"
	"lookingforpartner/common/params"
)

type OrderBy string

const (
	OrderByDESC OrderBy = "desc"
	OrderByASC  OrderBy = "asc"
)

const (
	CreatedAt = "created_at"
	UpdatedAt = "updated_at"
)

type OrderOpt struct {
	ColumnName string
	OrderBy    OrderBy
}

func OrderByString2Opt(o string) OrderOpt {
	switch o {
	case params.OrderByCreateTimeASC:
		return OrderOpt{
			ColumnName: CreatedAt,
			OrderBy:    OrderByASC,
		}
	case params.OrderByCreateTimeDESC:
		return OrderOpt{
			ColumnName: CreatedAt,
			OrderBy:    OrderByDESC,
		}
	case params.OrderByUpdateTimeASC:
		return OrderOpt{
			ColumnName: UpdatedAt,
			OrderBy:    OrderByASC,
		}
	case params.OrderByUpdateTimeDESC:
		return OrderOpt{
			ColumnName: UpdatedAt,
			OrderBy:    OrderByDESC,
		}
	default:
		return OrderOpt{
			ColumnName: CreatedAt,
			OrderBy:    OrderByDESC,
		}
	}
}

func (opt *OrderOpt) String() string {
	return fmt.Sprintf("%s %s", opt.ColumnName, opt.OrderBy)
}
