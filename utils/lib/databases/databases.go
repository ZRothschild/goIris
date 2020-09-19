package databases

import (
	"gorm.io/gorm"
	"math"
)

// Create
func Create(dB *gorm.DB, value interface{}) (rowsAffected int64, err error) {
	tx := dB.Create(value)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// Save
func Save(dB *gorm.DB, value interface{}) (rowsAffected int64, err error) {
	tx := dB.Save(value)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// Updates
func Updates(dB *gorm.DB, model, value, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	tx := dB.Model(model).Where(query, args...).Updates(value)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// Delete
func Delete(dB *gorm.DB, value, query interface{}, args ...interface{}) (rowsAffected int64, err error) {
	tx := dB.Where(query, args...).Delete(value)
	err = tx.Error
	rowsAffected = tx.RowsAffected
	return
}

// Delete
func DeleteById(dB *gorm.DB, value interface{}, id uint64) (rowsAffected int64, err error) {
	tx := dB.Delete(value, id)
	err = tx.Error
	rowsAffected = tx.RowsAffected
	return
}

// Delete
func DeleteByIds(dB *gorm.DB, value interface{}, ids []uint64) (rowsAffected int64, err error) {
	tx := dB.Delete(value, ids)
	err = tx.Error
	rowsAffected = tx.RowsAffected
	return
}

// First
func First(dB *gorm.DB, out interface{}, where interface{}, args ...interface{}) error {
	return dB.Where(where, args...).First(out).Error
}

// FirstWhere
func FirstWhere(dB *gorm.DB, out interface{}, query interface{}, wheres ...Condition) error {
	tx := dB
	if query != nil {
		tx = tx.Where(query)
	}
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				tx = tx.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				tx = tx.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				tx = tx.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				tx = tx.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				tx = tx.Order(where.Key)
			}
		}
	}
	return tx.First(out).Error
}

// FirstById
func FirstById(dB *gorm.DB, out interface{}, id uint64) error {
	return dB.First(out, id).Error
}

// FirstByIdWhere
func FirstByIdWhere(dB *gorm.DB, out interface{}, id uint64, wheres ...Condition) error {
	tx := dB
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				tx = tx.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				tx = tx.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				tx = tx.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				tx = tx.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				tx = tx.Order(where.Key)
			}
		}
	}
	return tx.First(out, id).Error
}

// Find
func Find(dB *gorm.DB, out interface{}, query interface{}, args ...interface{}) error {
	return dB.Where(query, args...).Find(out).Error
}

// FindWhere
func FindWhere(dB *gorm.DB, out interface{}, query interface{}, wheres ...Condition) error {
	tx := dB
	if query != nil {
		tx = tx.Where(query)
	}
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				tx = tx.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				tx = tx.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				tx = tx.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				tx = tx.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				tx = tx.Order(where.Key)
			}
		}
	}
	return tx.Find(out).Error
}

// FindByIds
func FindByIds(dB *gorm.DB, out interface{}, ids []uint64) error {
	return dB.Find(out, ids).Error
}

// FindByIdsWhere
func FindByIdsWhere(dB *gorm.DB, out interface{}, ids []uint64, wheres ...Condition) error {
	tx := dB
	if wheres != nil && len(wheres) > 0 {
		// where group  having  select order
		for _, where := range wheres {
			if where.Type == "select" && where.Key != nil {
				tx = tx.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				tx = tx.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				tx = tx.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				tx = tx.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				tx = tx.Order(where.Key)
			}
		}
	}
	return dB.Find(out, ids).Error
}

// GetPage
// 	var behaviourFront []model.BehaviourFront
// 	err := db.dB.Where("id_mock > ?",10).Select("req_uri,sum(id) as total").Group("req_uri,method").Having("count(id) > 7").Find(&behaviourFront)
func GetPage(dB *gorm.DB, model, query interface{}, out interface{}, pagination *Pagination, wheres ...Condition) error {
	tx := dB.Model(model)
	if query != nil {
		tx = tx.Where(query)
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
				tx = tx.Select(where.Key, where.Value...)
			} else if where.Type == "where" && where.Key != nil {
				tx = tx.Where(where.Key, where.Value...)
			} else if key, ok := where.Key.(string); ok && where.Type == "group" && where.Key != "" {
				tx = tx.Group(key)
			} else if where.Type == "having" && where.Key != nil {
				tx = tx.Having(where.Key, where.Value...)
			} else if where.Type == "order" && where.Key != nil {
				tx = tx.Order(where.Key)
			}
		}
	}
	if err := tx.Count(&pagination.Total).Error; err != nil {
		return err
	}

	if pagination.Total == 0 {
		return nil
	}

	if pagination.Page < 0 {
		return tx.Find(out).Error
	}
	// 总条数
	pagination.TotalPage = int(math.Ceil(float64(int(pagination.Total) / pagination.PageSize)))
	return tx.Offset((pagination.Page - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(out).Error
}

// Scan
func Scan(dB *gorm.DB, model, query, out interface{}, args ...interface{}) error {
	return dB.Model(model).Where(query, args...).Scan(out).Error
}

// ScanList
func ScanList(dB *gorm.DB, model, where interface{}, out interface{}, orders ...string) error {
	tx := dB.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			tx = tx.Order(order)
		}
	}
	return tx.Scan(out).Error
}

// PluckList
func PluckList(dB *gorm.DB, model, where, out interface{}, fieldName string, args ...interface{}) error {
	return dB.Model(model).Where(where, args...).Pluck(fieldName, out).Error
}
