package common

import (
	"fmt"
	"github.com/spf13/viper"
)

// 读取配置 使用viper包
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("读取配置：viper.AllKeys(): ", viper.AllKeys())
	fmt.Println("读取配置：config Mysql: ", viper.Get("Mysql"))
	fmt.Println("读取配置：config Mysql.dns: ", viper.GetString("Mysql.dns"))
}
