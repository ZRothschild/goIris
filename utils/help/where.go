package help

import (
	"github.com/ZRothschild/goIris/config/conf"
)

// 查询条件
func Where(field string, value ...interface{}) conf.Where {
	var where = conf.Where{
		Type: "where",
	}
	if field != "" {
		where.Key = field
	}

	if value != nil {
		where.Value = append(where.Value, value...)
	}
	return where
}

// 查询条件
func Condition(field interface{}, value ...interface{}) conf.Where {
	var where = conf.Where{
		Type: "where",
	}
	if field != nil {
		where.Key = field
	}

	if value != nil {
		where.Value = append(where.Value, value...)
	}
	return where
}

// BetweenInt64 条件
func BetweenInt64Where(field string, start, end int64) conf.Where {
	var where = conf.Where{
		Type: "where",
	}

	if start > 0 && end > start {
		where.Key = field + " >= ? and " + field + " <= ?"
		where.Value = append(where.Value, start, end)
	}
	return where
}

// in 条件
func InInt64Where(field string, ids []uint64) conf.Where {
	var where = conf.Where{
		Type: "where",
	}
	if field != "" && len(ids) > 0 {
		where.Key = field + " in (?)"
		where.Value = append(where.Value, ids)
	}
	return where
}

// 字段限制
func Select(field string, value ...interface{}) conf.Where {
	var where = conf.Where{
		Type: "select",
	}
	if field != "" {
		where.Key = field

	}
	if value != nil {
		where.Value = append(where.Value, value...)
	}
	return where
}

// 排序
func OrderBy(field string, sortType string) conf.Where {
	var where = conf.Where{
		Type: "order",
	}
	if field != "" && sortType != "" {
		where.Key = field + " " + sortType
	}
	return where
}

// 多字段排序
func OrderBys(fields []string, sortTypes []string) conf.Where {
	var (
		sql   string
		where = conf.Where{
			Type: "order",
		}
		fieldLen    = len(fields)
		sortTypeLen = len(sortTypes)
	)
	if fieldLen != sortTypeLen {
		return where
	}

	for k, v := range fields {
		if k == 0 {
			sql = v + " " + sortTypes[k]
		} else {
			sql += "," + v + " " + sortTypes[k]
		}
	}
	where.Key = sql
	return where
}
