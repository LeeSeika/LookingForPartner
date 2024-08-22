package params

import "lookingforpartner/common/dao"

var (
	OrderByCreateTimeASC  = "create_time_asc"
	OrderByCreateTimeDESC = "create_time_desc"
	OrderByUpdateTimeASC  = "update_time_asc"
	OrderByUpdateTimeDESC = "update_time_desc"
)

func ToOrderByOpt(o string) dao.OrderOpt {
	switch o {
	case OrderByCreateTimeASC:
		return dao.OrderOpt{
			ColumnName: dao.CreatedAt,
			OrderBy:    dao.OrderByASC,
		}
	case OrderByCreateTimeDESC:
		return dao.OrderOpt{
			ColumnName: dao.CreatedAt,
			OrderBy:    dao.OrderByDESC,
		}
	case OrderByUpdateTimeASC:
		return dao.OrderOpt{
			ColumnName: dao.UpdatedAt,
			OrderBy:    dao.OrderByASC,
		}
	case OrderByUpdateTimeDESC:
		return dao.OrderOpt{
			ColumnName: dao.UpdatedAt,
			OrderBy:    dao.OrderByDESC,
		}
	default:
		return dao.OrderOpt{
			ColumnName: dao.CreatedAt,
			OrderBy:    dao.OrderByDESC,
		}
	}
}
