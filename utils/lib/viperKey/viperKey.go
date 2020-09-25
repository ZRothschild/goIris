package viperKey

import (
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

// 获取配置前缀
func PreKeyViper(newViper *viper.Viper, keyStr ...string) (key string, err error) {
	var (
		global = new(Global)
	)
	if err = newViper.UnmarshalKey(Glob, global); err != nil {
		return key, err
	}

	key = global.App + "." + global.Env + "."

	length := len(keyStr)
	for k, v := range keyStr {
		if length == k+1 {
			key += v
			return key, err
		}
		key += v + "."
	}

	return key, err
}
