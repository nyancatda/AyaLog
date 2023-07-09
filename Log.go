/*
 * @Author: NyanCatda
 * @Date: 2022-05-22 00:03:28
 * @LastEditTime: 2023-01-07 22:02:49
 * @LastEditors: NyanCatda
 * @Description: 日志模块
 * @FilePath: \AyaLog\Log.go
 */
package AyaLog

import "os"

// 定义日志等级
const (
	DEBUG = iota + 0
	INFO
	WARNING
	ERROR
	OFF // 关闭日志
)

type Log struct {
	Path            string // 日志文件保存路径
	Segmentation    string // 日志文件分割标识(使用go默认时间格式)
	WriteFile       bool   // 是否写入文件
	ColorPrint      bool   // 是否打印颜色
	Level           int    // 日志等级
	Prefix          string // 日志前缀
	PrefixWriteFile bool   // 日志前缀是否写入文件
	Suffix          string // 日志后缀
	SuffixWriteFile bool   // 日志后缀是否写入文件
	PrintErrorStack bool   // 是否打印错误堆栈

	nowSegmentation string   // 当前日志分割标识
	logFile         *os.File // 日志文件实例

	middlewareBefore []middlewareBefore // 日志打印前中间件组
	middlewareAfter  []middlewareAfter  // 日志打印后中间件组
}

type LogPrint interface {
	Print(Source string, Level int, Text ...any) error // 打印日志
	Error(Source string, Error error, Text ...any)     // 打印错误
	Warning(Source string, Text ...any)                // 打印警告
	Info(Source string, Text ...any)                   // 打印信息
	DeBug(Source string, Text ...any)                  // 打印DeBug信息
}

/**
 * @description: 创建一个默认日志实例
 * @param {*}
 * @return {Log} 日志实例
 */
func NewLog() *Log {
	return &Log{
		Path:            "./logs/",
		Segmentation:    "2006-01-02",
		WriteFile:       true,
		ColorPrint:      true,
		Level:           DEBUG,
		Prefix:          "",
		PrefixWriteFile: false,
		Suffix:          "",
		SuffixWriteFile: false,
		PrintErrorStack: true,
	}
}
