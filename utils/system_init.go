package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"qqchat/common"
	"qqchat/utils/qqlog"
	"time"
)

var Db *gorm.DB

func init() {
	InitConfig()           // 初始化配置
	qqlog.InitLogger()     //初始化日志Logrus日志库
	common.InitValidator() // 初始化  自定义校验器 【暂时没用到】
	InitGorm()             //初始化gorm
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

// InitGorm gorm初始化 数据库初始化 并写入慢日志
func InitGorm() *gorm.DB {
	// mysql 日志配置
	log_dir := viper.GetString("Mysql.LOG_DIR")
	// 获取当前项目跟目录路径
	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatal("获取当前项目跟目录路径:", err)
	}

	// 创建日志目录
	logDir := filepath.Join(projectDir, log_dir)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("创建日志目录失败:", err)
	}
	fmt.Println("mysql日志目录路径为：", logDir)

	// 生成带日期的日志文件名
	now := time.Now()
	date := now.Format("20060102") // YYYYMMDD格式
	sqlLogFile := filepath.Join(logDir, fmt.Sprintf("slow_sql_%s.log", date))
	fmt.Println("mysql日志具体路径为：", sqlLogFile)

	// 初始化日志文件
	var sqlLog *os.File
	sqlLog, _ = os.OpenFile(
		sqlLogFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0755)

	// sql 日志
	newLogger := logger.New(log.New(sqlLog, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 2 * time.Second, // 超过2秒的SQL视为慢查询
		LogLevel:      logger.Info,     // 记录INFO级别及以上日志
		Colorful:      true,            // 启用彩色日志输出
	})

	// mysql 连接
	// 创建SQL日志记录器
	// 使用标准输出流记录日志，配置日志级别和格式
	db, err := gorm.Open(mysql.Open(viper.GetString("Mysql.dns")), &gorm.Config{
		Logger: newLogger, // 注入自定义日志记录器
	})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	// 将数据库连接对象挂载到全局变量Db
	Db = db
	return Db
}
