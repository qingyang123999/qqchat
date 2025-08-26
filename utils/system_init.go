package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	InitConfig()
	InitGorm()
}

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

// InitGorm gorm初始化 数据库初始化
func InitGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(viper.GetString("Mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	Db = db
	return Db
}
