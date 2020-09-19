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
	App string
	Env string
}

/**
获取数据库配置文件viper key
key 获取类型 MySql配置
*/
func MySql(key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey("MySql", newViper)
	return str + key, err
}

/**
获取数据库配置文件viper key
key 获取类型 Log配置
*/
func Log(key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey("Log", newViper)
	return str + key, err
}

/**
获取数据库配置文件viper key
key 获取类型 validators 配置
*/
func Validator(key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey("Validator", newViper)
	return str + key, err
}

// 获取配置前缀
func preViperKey(keyType string, newViper *viper.Viper) (string, error) {
	var (
		key    string
		err    error
		global = new(Global)
	)
	if err = newViper.UnmarshalKey(Glob, global); err != nil {
		return key, err
	}

	key = global.App + "." + global.Env + "." + keyType + "."
	return key, errors.New("配置文件错误")
}
