// Package date 日期工具类
package date

import "time"

// CurrentDateTime 获取当前日期及时间，格式：2006-01-02 15:04:05
func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurrentDate 获取当前日期，格式：2006-01-02
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// CurrentTime 获取当前时间，格式：15:04:05
func CurrentTime() string {
	return time.Now().Format("15:04:05")
}

// FormatDate 自定义格式化当前时间
func FormatDate(layout string) string {
	return time.Now().Format(layout)
}

// Convert 日期类型转换
// @date 日期字符串
// @srcFormat 源日期格式
// @dstFormat 目标日期格式
func Convert(date string, srcFormat string, dstFormat string) string {
	srcDate, err := time.Parse(srcFormat, date)
	if err != nil {
		return ""
	}
	return srcDate.Format(dstFormat)
}
