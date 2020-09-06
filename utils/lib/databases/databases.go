package databases

import (
	"gorm.io/gorm"
	"math"
)

// Create
func Create(dB *gorm.DB, value interface{}) (rowsAffected int64, err error) {
	dB = dB.Create(value)
	rowsAffected = dB.RowsAffected
	err = dB.Error
	return
}

// Save
func Save(dB *gorm.DB, value interface{}) (rowsAffected int64, err error) {
	dB = dB.Save(value)
	rowsAffected = dB.RowsAffected
	err = dB.Error
	return
}

// Updates
func Updates(dB *gorm.DB, model, value, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	dB = dB.Model(model).Where(query, args...).Updates(value)
	rowsAffected = dB.RowsAffected
	err = dB.Error
	return
}

// Delete
func Delete(dB *gorm.DB, value, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	dB = dB.Where(query, args...).Delete(value)
	err = dB.Error
	rowsAffected = dB.RowsAffected
	return
}

// Delete
func DeleteById(dB *gorm.DB, value interface{}, id uint64) (rowsAffected int64, err error) {
	dB = dB.Delete(value, id)
	err = dB.Error
	rowsAffected = dB.RowsAffected
	return
}

// Delete
func DeleteByIds(dB *gorm.DB, value interface{}, ids []uint64) (rowsAffected int64, err error) {
	dB = dB.Delete(value, ids)
	err = dB.Error
	rowsAffected = dB.RowsAffected
	return
}

// First
func First(dB *gorm.DB, where interface{}, out interface{}, args ...interface{}) error {
	return dB.Where(where, args...).First(out).Error
}

// FirstWhere
func FirstWhere(dB *gorm.DB, out interface{}, query interface{}, wheres ...Condition) error {
	if query != nil {
		dB = dB.Where(query)
	}
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				dB = dB.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				dB = dB.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				dB = dB.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				dB = dB.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				dB = dB.Order(where.Key)
			}
		}
	}
	return dB.First(out).Error
}

// FirstById
func FirstById(dB *gorm.DB, out interface{}, id uint64) error {
	return dB.First(out, id).Error
}

// FirstByIdWhere
func FirstByIdWhere(dB *gorm.DB, out interface{}, id uint64, wheres ...Condition) error {
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				dB = dB.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				dB = dB.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				dB = dB.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				dB = dB.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				dB = dB.Order(where.Key)
			}
		}
	}
	return dB.First(out, id).Error
}

// Find
func Find(dB *gorm.DB, out interface{}, query interface{}, args ...interface{}) error {
	return dB.Where(query, args...).Find(out).Error
}

// FindWhere
func FindWhere(dB *gorm.DB, out interface{}, query interface{}, wheres ...Condition) error {
	if query != nil {
		dB = dB.Where(query)
	}
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				dB = dB.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				dB = dB.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				dB = dB.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				dB = dB.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				dB = dB.Order(where.Key)
			}
		}
	}
	return dB.Find(out).Error
}

// FindByIds
func FindByIds(dB *gorm.DB, out interface{}, ids []uint64) error {
	return dB.Find(out, ids).Error
}

// FindByIdsWhere
func FindByIdsWhere(dB *gorm.DB, out interface{}, ids []uint64, wheres ...Condition) error {
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				dB = dB.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				dB = dB.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				dB = dB.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				dB = dB.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				dB = dB.Order(where.Key)
			}
		}
	}
	return dB.Find(out, ids).Error
}

// GetPage
// 	var behaviourFront []model.BehaviourFront
// 	err := db.dB.Where("id_mock > ?",10).Select("req_uri,sum(id) as total").Group("req_uri,method").Having("count(id) > 7").Find(&behaviourFront)
func GetPage(dB *gorm.DB, model, query interface{}, out interface{}, pagination *Pagination, wheres ...Condition) error {
	dB = dB.Model(model)
	if query != nil {
		dB = dB.Where(query)
	}
	// 默认第一页 每页二十条
	if pagination.Page == 0 {
		pagination.Page = 1
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 20
	}
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				dB = dB.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				dB = dB.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				dB = dB.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				dB = dB.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				dB = dB.Order(where.Key)
			}
		}
	}
	if err := dB.Count(&pagination.Total).Error; err != nil {
		return err
	}

	if pagination.Total == 0 {
		return nil
	}

	if pagination.Page < 0 {
		return dB.Find(out).Error
	}
	// 总条数
	pagination.TotalPage = int(math.Ceil(float64(int(pagination.Total) / pagination.PageSize)))
	return dB.Offset((pagination.Page - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(out).Error
}

// Scan
func Scan(dB *gorm.DB, model, query, out interface{}, args ...interface{}) error {
	return dB.Model(model).Where(query, args...).Scan(out).Error
}

// ScanList
func ScanList(dB *gorm.DB, model, where interface{}, out interface{}, orders ...string) error {
	tDb := dB.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			tDb = tDb.Order(order)
		}
	}
	return tDb.Scan(out).Error
}

// PluckList
func PluckList(dB *gorm.DB, model, where, out interface{}, fieldName string, args ...interface{}) error {
	return dB.Model(model).Where(where, args...).Pluck(fieldName, out).Error
}
