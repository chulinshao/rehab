package util

import "time"

// 获取当前时间字符串
func CurrDateStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
