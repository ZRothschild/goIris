package db

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/utils/help/id"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

// 生成链接对象
func NewMySql(viper *viper.Viper, viperKey string) (*gorm.DB, error) {
	// 初始化数据库
	var (
		dns Dns
		err error
		dB  *gorm.DB
	)
	if err = viper.UnmarshalKey(viperKey, &dns); err != nil {
		return dB, err
	}
	dnsStr := dns.User + ":" + dns.Pwd + "@tcp(" + dns.Host + ":" + dns.Port + ")/" + dns.DbName + "?charset=utf8&parseTime=True&loc=Local"
	dB, err = gorm.Open(mysql.Open(dnsStr), &gorm.Config{})
	if err != nil {
		return dB, err
	}

	if err = dB.Callback().Create().Before("gorm:create").Register("fill_id", func(db *gorm.DB) {
		// 自动填充id
		db.Statement.SetColumn("id", id.NextId(viper))
	}); err != nil {
		return dB, err
	}

	sqlDB, err := dB.DB()
	if err != nil {
		return dB, err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	dB = dB.Debug()

	// 数据迁移
	err = migration(dB)
	return dB, err
}

// 数据迁移
func migration(dB *gorm.DB) (err error) {
	// 传入要迁移的模型指针
	return dB.AutoMigrate(
		new(model.User),
	)
}
