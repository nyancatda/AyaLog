/*
 * @Author: NyanCatda
 * @Date: 2022-05-22 00:03:28
 * @LastEditTime: 2022-07-19 18:23:17
 * @LastEditors: NyanCatda
 * @Description: 日志模块
 * @FilePath: \AyaLog\Log.go
 */
package AyaLog

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	LogPath         = "./logs/"    // 日志文件保存路径
	LogSegmentation = "2006-01-02" // 日志文件分割标识
	LogWriteFile    = true         // 是否写入文件
	ColorPrint      = true         // 是否打印颜色
	LogLevel        = DEBUG        // 日志等级
	LineEndString   = ""           // 行末字符串(不会被打印到日志文件)
)

// 定义日志等级
const (
	DEBUG = iota + 0
	INFO
	WARNING
	ERROR
	OFF // 关闭日志
)

/**
 * @description: 打印错误
 * @param {string} Source 日志来源
 * @param {error} Error 错误信息
 * @return {*}
 */
func Error(Source string, Error error) {
	// 追踪错误来源
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	ErrorStack := fmt.Sprintf("\n%s", string(buf[:n]))

	Print(Source, ERROR, Error.Error()+ErrorStack)
}

/**
 * @description: 打印警告
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func Warning(Source string, Text ...any) {
	Print(Source, WARNING, Text...)
}

/**
 * @description: 打印信息
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func Info(Source string, Text ...any) {
	Print(Source, INFO, Text...)
}

/**
 * @description: 打印DeBug错误
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func DeBug(Source string, Text ...any) {
	Print(Source, DEBUG, Text...)
}

/**
 * @description:  标准日志打印
 * @param {string} Source 日志来源
 * @param {string} Level 日志等级 INFO/WARNING/ERROR/DEBUG
 * @param {...any} Text 日志内容
 * @return {*}
 */
func Print(Source string, Level int, Text ...any) error {
	// 根据日志等级判断是否打印
	if Level < LogLevel {
		return nil
	}

	// 等级打印OFF则不打印
	if Level >= OFF {
		return nil
	}

	// 获取当前时间
	NowTime := time.Now().Format("2006-01-02 15:04:05")

	// Source拼接
	Source = "[" + Source + "]"

	// 判断level颜色
	var LevelStr string
	switch Level {
	case 0:
		LevelStr = Green("DEBUG")
	case 1:
		LevelStr = Blue("INFO")
	case 2:
		LevelStr = Yellow("WARNING")
	case 3:
		LevelStr = Red("ERROR")
	}

	Text = append([]any{Cyan(NowTime), LevelStr, Source}, Text...)

	// 准备打印日志
	var LogText []any
	// 如果彩色打印被关闭
	if !ColorPrint {
		// 遍历消息内容去除颜色
		for _, v := range Text {
			DelColorText := DelColor(fmt.Sprint(v))
			LogText = append(LogText, DelColorText)
		}
	} else {
		LogText = Text
	}

	if LineEndString != "" {
		// 添加行末字符串
		LogText = append(LogText, LineEndString)
	}

	// 打印日志
	_, err := fmt.Println(LogText...)
	if err != nil {
		return err
	}

	// 写入日志
	if LogWriteFile {
		logFile, err := LogFile()
		if err != nil {
			fmt.Println(err)
		}
		defer logFile.Close()
		write := bufio.NewWriter(logFile)

		// 遍历消息内容去除颜色
		var LogFileText string
		for _, v := range Text {
			DelColorText := DelColor(fmt.Sprint(v))
			LogFileText += DelColorText
			LogFileText += " "
		}

		write.WriteString(LogFileText + "\n")
		write.Flush()
	}

	return nil
}

/**
 * @description: 打开Log文件，按天分割日志
 * @param {*}
 * @return {*os.File}
 * @return {error}
 */
func LogFile() (*os.File, error) {
	// 判断文件夹是否存在
	MKDir(LogPath)

	logFileName := time.Now().Format(LogSegmentation) + ".log"

	logfile, err := os.OpenFile(LogPath+logFileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果文件不存在则创建
		logfile, err := os.Create(LogPath + logFileName)
		if err != nil {
			return logfile, err
		}
		return logfile, nil
	}

	return logfile, nil
}

/**
 * @description: 创建文件夹，如果不存在则创建
 * @param {string} path 文件夹路径
 * @return {*}
 */
func MKDir(Path string) (bool, error) {
	Path = filepath.Clean(Path)
	_, err := os.Stat(Path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
