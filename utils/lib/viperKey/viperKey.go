package viperKey

import (
	"errors"
	"github.com/spf13/viper"
)

// 需要参考config配置文件
const (
	Glob = "Global"
)

// config 配置文件 Global 结构体
type Global struct {
	App      string
	Env      string
	Backend  string
	Frontend string
}

/**
获取数据库配置文件viper key
source 前台或后台
key 获取类型 MySql配置
*/
func MySql(source, key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey(source, "MySql", newViper)
	return str + key, err
}

/**
获取数据库配置文件viper key
source 前台或后台
key 获取类型 Log配置
*/
func Log(source, key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey(source, "Log", newViper)
	return str + key, err
}

// 获取配置前缀
func preViperKey(source, keyType string, newViper *viper.Viper) (string, error) {
	var (
		key    string
		err    error
		global = new(Global)
	)
	if err = newViper.UnmarshalKey(Glob, global); err != nil {
		return key, err
	}
	if source == global.Backend {
		key = global.App + "." + global.Backend + "." + global.Env + "." + keyType + "."
		return key, err
	}

	if source == global.Frontend {
		key = global.App + "." + global.Frontend + "." + global.Env + "." + keyType + "."
		return key, err
	}
	return key, errors.New("配置文件错误")
}
