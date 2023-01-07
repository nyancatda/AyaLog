/*
 * @Author: NyanCatda
 * @Date: 2022-11-26 16:50:36
 * @LastEditTime: 2023-01-07 22:58:09
 * @LastEditors: NyanCatda
 * @Description: 打印日志
 * @FilePath: \AyaLog\Print.go
 */
package AyaLog

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/**
 * @description: 打印错误
 * @param {string} Source 日志来源
 * @param {error} Error 错误信息
 * @return {*}
 */
func (Log *Log) Error(Source string, Error error) {
	if Log.PrintErrorStack {
		// 追踪错误来源
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		ErrorStack := fmt.Sprintf("\n%s", string(buf[:n]))

		Log.Print(Source, ERROR, Error.Error(), ErrorStack)
		return
	}

	Log.Print(Source, ERROR, Error)
}

/**
 * @description: 打印警告
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func (Log *Log) Warning(Source string, Text ...any) {
	Log.Print(Source, WARNING, Text...)
}

/**
 * @description: 打印信息
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func (Log *Log) Info(Source string, Text ...any) {
	Log.Print(Source, INFO, Text...)
}

/**
 * @description: 打印DeBug错误
 * @param {string} Source 日志来源
 * @param {...any} Text 日志内容
 * @return {*}
 */
func (Log *Log) DeBug(Source string, Text ...any) {
	Log.Print(Source, DEBUG, Text...)
}

/**
 * @description:  标准日志打印
 * @param {string} Source 日志来源
 * @param {int} Level 日志等级 DEBUG/INFO/WARNING/ERROR/OFF
 * @param {...any} Text 日志内容
 * @return {error} error
 */
func (Log *Log) Print(Source string, Level int, Text ...any) error {
	// 执行日志打印前中间件
	if len(Log.middlewareBefore) != 0 {
		var TextPointer []*any
		for i := range Text {
			TextPointer = append(TextPointer, &Text[i])
		}

		for _, Func := range Log.middlewareBefore {
			Func(&Level, &Source, TextPointer...)
		}
	}

	// 根据日志等级判断是否打印
	if Level < Log.Level {
		return nil
	}

	// 等级打印OFF则不打印
	if Level >= OFF {
		return nil
	}

	// 获取当前时间
	NowTime := time.Now().Format("2006-01-02 15:04:05")

	// Source拼接
	PrintSource := "[" + Source + "]"

	// 判断level颜色
	var PrintLevel string
	switch Level {
	case 0:
		PrintLevel = Green("DEBUG")
	case 1:
		PrintLevel = Blue("INFO")
	case 2:
		PrintLevel = Yellow("WARNING")
	case 3:
		PrintLevel = Red("ERROR")
	default:
		PrintLevel = strconv.Itoa(Level)
	}

	var PrintText []any
	PrintText = append([]any{Cyan(NowTime), PrintLevel, PrintSource}, Text...)

	// 准备打印日志
	var LogBody []any
	// 如果彩色打印被关闭
	if !Log.ColorPrint {
		// 遍历消息内容去除颜色
		for _, v := range PrintText {
			DelColorText := DelColor(fmt.Sprint(v))
			LogBody = append(LogBody, DelColorText)
		}
	} else {
		LogBody = PrintText
	}

	// 追加打印前缀
	if Log.Prefix != "" {
		LogBody = append([]any{Log.Prefix}, LogBody...)
	}
	// 追加打印后缀
	if Log.Suffix != "" {
		LogBody = append(LogBody, Log.Suffix)
	}

	// 打印日志
	_, err := fmt.Println(LogBody...)
	if err != nil {
		return err
	}

	// 写入日志
	if Log.WriteFile {
		err := Log.writeLogFile(PrintText...)
		if err != nil {
			return err
		}
	}

	// 执行日志打印后中间件
	if len(Log.middlewareAfter) != 0 {
		for _, Func := range Log.middlewareAfter {
			Func(Level, Source, Text...)
		}
	}

	return nil
}
