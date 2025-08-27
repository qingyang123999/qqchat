package utils

import "time"

func GetNowTime() string {
	currentTime := time.Now() // 获取当前时间对象
	return currentTime.Format("2006-01-02 15:04:05")
}

func GetNowUnixTime() int64 {
	currentTime := time.Now() // 获取当前时间对象
	return currentTime.Unix()
}
