package help

import (
	"errors"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/spf13/viper"
)

// 获取数据库配置文件viper key
func FrontMySqlViperKey(key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey(conf.Frontend, "MySql", newViper)
	return str + key, err
}

func FrontLogViperKey(key string, newViper *viper.Viper) (string, error) {
	var (
		str string
		err error
	)
	str, err = preViperKey(conf.Frontend, "Log", newViper)
	return str + key, err
}

// 获取配置前缀
func preViperKey(team, keyType string, newViper *viper.Viper) (string, error) {
	var (
		key    string
		err    error
		global = new(conf.Global)
	)
	if err = newViper.UnmarshalKey(conf.Glob, global); err != nil {
		return key, err
	}
	if team == global.Backend {
		key = global.App + "." + global.Backend + "." + global.Env + "." + keyType + "."
		return key, err
	}

	if team == global.Frontend {
		key = global.App + "." + global.Frontend + "." + global.Env + "." + keyType + "."
		return key, err
	}
	return key, errors.New("配置文件错误")
}
