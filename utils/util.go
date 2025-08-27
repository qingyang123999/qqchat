package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func GetNowTime() string {
	currentTime := time.Now() // 获取当前时间对象
	return currentTime.Format("2006-01-02 15:04:05")
}

func GetNowUnixTime() int64 {
	currentTime := time.Now() // 获取当前时间对象
	return currentTime.Unix()
}

// 在项目下创建新文件，并打开新文件，关闭旧文件
func MakeDoc() {
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
	today := now.Format("20060102") // YYYYMMDD格式

	// 检查是否需要切换日志文件
	var sqlLog *os.File
	var lastDate string

	// 尝试从日志中提取最后使用的日期
	if sqlLog != nil {
		_, err := sqlLog.Seek(0, io.SeekEnd)
		if err == nil {
			// 从路径中提取日期
			lastDate = filepath.Base(sqlLog.Name())
			if len(lastDate) >= 8 {
				lastDate = lastDate[4:12] // 提取日期部分
			}
		}
	}

	// 如果日期变化，则关闭旧文件并打开新文件
	if lastDate != today {
		if sqlLog != nil {
			sqlLog.Close()
		}

		sqlLogFile := filepath.Join(logDir, fmt.Sprintf("sql_%s.log", today))
		sqlLog, err = os.OpenFile(
			sqlLogFile,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644)
		if err != nil {
			log.Fatal("打开SQL日志文件失败:", err)
		}
	}
}
