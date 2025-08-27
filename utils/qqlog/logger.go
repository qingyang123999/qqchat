package qqlog

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// 使用示例
//// 初始化日志
//qlog.InitLogger()
//
//// 基础日志记录
//qlog.Log.Info("服务启动成功")
//qlog.Log.WithFields(logrus.Fields{
//"user": "admin",
//"ip":   "192.168.1.1",
//}).Warn("异常登录尝试")
//
//// Gin框架使用
//r := gin.New()
//r.Use(qlog.LoggerMiddleware())

var Log *logrus.Logger

// 日志级别(根据环境自动设置)  当检测到时生产环境时。只写入error和致命错误级别的日志。不写入其它级别的日志
func InitLogger() {
	Log = logrus.New()

	// 基础配置
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
	})

	// 日志级别(可从环境变量读取)
	Log.SetLevel(logrus.InfoLevel)

	// 日志文件Hook配置
	logDir := viper.GetString("Logger.LOG_DIR")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Errorf("创建日志目录失败: %s \n", err))
	}
	fmt.Println("logger日志目录路径为：", logDir)

	// 日志级别(根据环境自动设置)
	env := os.Getenv("APP_ENV")
	switch {
	case env == "prod" || env == "production":
		Log.SetLevel(logrus.ErrorLevel) // 生产环境只记录错误级别以上
	default:
		Log.SetLevel(logrus.InfoLevel) // 其他环境保持原配置
	}

	// 动态调整writerMap
	writerMap := lfshook.WriterMap{}
	if env == "prod" || env == "production" {
		// 生产环境只配置error级别
		writerMap[logrus.ErrorLevel] = createLogFileWriter("error", logDir)
	} else {
		// 其他环境保持完整配置
		writerMap = lfshook.WriterMap{
			logrus.InfoLevel:  createLogFileWriter("info", logDir),
			logrus.DebugLevel: createLogFileWriter("debug", logDir),
			logrus.WarnLevel:  createLogFileWriter("warn", logDir),
			logrus.ErrorLevel: createLogFileWriter("error", logDir),
		}
	}

	Log.AddHook(lfshook.NewHook(
		writerMap,
		&logrus.JSONFormatter{},
	))

	// 添加自定义字段(全局)
	Log.AddHook(&DefaultFieldHook{
		AppName: "myapp",
		Env:     os.Getenv("APP_ENV"),
	})
}

// createLogFileWriter 创建目录后再 创建支持日志切割的文件写入器
func createLogFileWriter(level string, logDir string) *os.File {
	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Dir(logDir), 0755); err != nil {
		logrus.WithError(err).Error("创建日志目录失败")
		panic(fmt.Errorf("创建日志目录失败: %s \n", err))
		return os.Stderr
	}

	// 打开或创建日志文件
	file, err := os.OpenFile(
		getCurrentLogFile(level, logDir),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		logrus.WithError(err).Error("打开日志文件失败")
		panic(fmt.Errorf("打开日志文件失败: %s \n", err))
		return os.Stderr
	}
	return file
}

// getCurrentLogFile 根据日期模板获取实际日志文件名
func getCurrentLogFile(level string, logDir string) string {
	// 生成带日期的日志文件名
	now := time.Now()
	date := now.Format("20060102") // YYYYMMDD格式
	logFilePath := filepath.Join(logDir, fmt.Sprintf("service_%s_%s.log", level, date))
	fmt.Println("mysql日志具体路径为：", logFilePath)

	return logFilePath
}
