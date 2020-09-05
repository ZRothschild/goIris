package viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	viper3 "github.com/spf13/viper"
)

func NewViper(confName,confType string,confPath ... string) (viper *viper3.Viper) {
	// 读取 yaml 配置文件 设置配置文件名为 config, 不需要配置文件扩展名，
	// 配置文件的类型 viper 会自动根据扩展名自动匹配.
	viper = viper3.New()
	viper.SetConfigName(confName)   // 配置文件名称
	if confPath == nil {
		viper.AddConfigPath("./conf/")
	}else {
		for _, path := range confPath  {
			viper.AddConfigPath(path)   // 配置文件名称
		}
	}
	viper.SetConfigType(confType)

	// viper.SetConfigName("config")   // 配置文件名称
	// viper.AddConfigPath("./conf/")  // 设置配置文件的搜索目录
	// viper.AddConfigPath("../conf/") // 设置配置文件的搜索目录
	// viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil { // 加载配置文件内容
		panic(fmt.Errorf("Fatal error viper config file: %s \n", err))
	}
	go ReloadConfig(viper)
	return
}

// 热加载
func ReloadConfig(viper *viper3.Viper) {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("viper Detect config change: %s \n", in.String())
	})
}
