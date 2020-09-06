package databases

type (
	// 页码结构体
	Pagination struct {
		Page      int   `json:"page" example:"0"`      // 当前页
		PageSize  int   `json:"pageSize" example:"20"` // 每页条数
		TotalPage int   `json:"totalPage"`             // 总页数
		Total     int64 `json:"total"`                 // 总条数
	}

	// 查询条件
	Condition struct {
		Type  string        // where group having select order
		Key   interface{}   // 表达式
		Value []interface{} // 值
	}
)

// 查询条件
func Where(field string, value ...interface{}) Condition {
	var where = Condition{
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
func Query(field interface{}, value ...interface{}) Condition {
	var where = Condition{
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
func BetweenInt64Where(field string, start, end int64) Condition {
	var where = Condition{
		Type: "where",
	}

	if start > 0 && end > start {
		where.Key = field + " >= ? and " + field + " <= ?"
		where.Value = append(where.Value, start, end)
	}
	return where
}

// in 条件
func InInt64Where(field string, ids []uint64) Condition {
	var where = Condition{
		Type: "where",
	}
	if field != "" && len(ids) > 0 {
		where.Key = field + " in (?)"
		where.Value = append(where.Value, ids)
	}
	return where
}

// 字段限制
func Select(field string, value ...interface{}) Condition {
	var where = Condition{
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
func OrderBy(field string, sortType string) Condition {
	var where = Condition{
		Type: "order",
	}
	if field != "" && sortType != "" {
		where.Key = field + " " + sortType
	}
	return where
}

// 多字段排序
func OrderBys(fields []string, sortTypes []string) Condition {
	var (
		sql   string
		where = Condition{
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
