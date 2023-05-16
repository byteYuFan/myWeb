package ttviper

import (
	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
}

func ConfigInit(cfgFileName string) Config {
	v := viper.New()
	config := Config{Viper: v}
	viper := config.Viper
	viper.SetConfigName(cfgFileName) // 设置配置文件名字
	viper.AddConfigPath(".")         // 添加配置文件路径，可以添加多个路径
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../../config")
	viper.SetConfigType("yml") // 配置文件类型
	err := config.Viper.ReadInConfig()
	if err != nil {
		panic(err)
		return Config{}
	}
	return config

}
