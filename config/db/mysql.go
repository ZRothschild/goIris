package db

import (
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math"
)

type (
	// 数据库链接
	Dns struct {
		Host   string
		Port   string
		Pwd    string
		User   string
		DbName string
	}
)

func NewMySql(viperKey string, viper *viper.Viper) *gorm.DB {
	// 初始化数据库
	var (
		dns Dns
		err error
		dB  *gorm.DB
	)
	if err := viper.UnmarshalKey(viperKey, &dns); err != nil {
		log.Fatal("数据库获取配置文件失败" + err.Error())
	}
	dnsStr := dns.User + ":" + dns.Pwd + "@tcp(" + dns.Host + ":" + dns.Port + ")/" + dns.DbName + "?charset=utf8&parseTime=True&loc=Local"
	dB, err = gorm.Open(mysql.Open(dnsStr), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库链接失败" + err.Error())
	}
	dB = dB.Debug()

	// 数据迁移
	migration(dB)
	return dB
}

// 数据迁移
func migration(dB *gorm.DB) {
	// 传入要迁移的模型指针
	if err := dB.AutoMigrate().Error; err != nil {
		// 错误
	}
}

// Create
func Create(dB *gorm.DB, value interface{}) error {
	return dB.Create(value).Error
}

// Save
func Save(dB *gorm.DB, value interface{}) error {
	return dB.Save(value).Error
}

// Updates
func Updates(dB *gorm.DB, model, value, query interface{}, args ...interface{}) error {
	return dB.Model(model).Where(query, args...).Updates(value).Error
}

// Delete
func Delete(dB *gorm.DB, query, model interface{}, args ...interface{}) (count int64, err error) {
	dB = dB.Where(query, args...).Delete(model)
	err = dB.Error
	if err != nil {
		return
	}
	count = dB.RowsAffected
	return
}

// Delete
func DeleteById(dB *gorm.DB, model interface{}, id uint64) (count int64, err error) {
	dB = dB.Delete(model, id)
	err = dB.Error
	if err != nil {
		return
	}
	count = dB.RowsAffected
	return
}

// Delete
func DeleteByIds(dB *gorm.DB, model interface{}, ids []uint64) (count int64, err error) {
	dB = dB.Delete(model, ids)
	err = dB.Error
	if err != nil {
		return
	}
	count = dB.RowsAffected
	return
}

// First
func FirstById(dB *gorm.DB, out interface{}, id uint64) error {
	return dB.First(out, id).Error
}

func FirstByIdWhere(dB *gorm.DB, out interface{}, id uint64, wheres ...conf.Where) error {
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

// First
func First(dB *gorm.DB, where interface{}, out interface{}, args ...interface{}) error {
	return dB.Where(where, args...).First(out).Error
}

// FirstWhere
func FirstWhere(dB *gorm.DB, out interface{}, query interface{}, wheres ...conf.Where) error {
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

// Find
func Find(dB *gorm.DB, out interface{}, query interface{}, args ...interface{}) error {
	return dB.Where(query, args...).Find(out).Error
}

// FindByIdsWhere
func FindByIdsWhere(dB *gorm.DB, out interface{}, ids []uint64, wheres ...conf.Where) error {
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

// FindWhere
func FindWhere(dB *gorm.DB, out interface{}, query interface{}, wheres ...conf.Where) error {
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

// GetPage
// 	var behaviourFront []model.BehaviourFront
// 	err := db.dB.Where("id_mock > ?",10).Select("req_uri,sum(id) as total").Group("req_uri,method").Having("count(id) > 7").Find(&behaviourFront)
func GetPage(dB *gorm.DB, model, query interface{}, out interface{}, pagination *conf.Pagination, wheres ...conf.Where) error {
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
