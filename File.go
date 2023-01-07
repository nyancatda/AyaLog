/*
 * @Author: NyanCatda
 * @Date: 2022-11-26 16:45:50
 * @LastEditTime: 2023-01-07 17:22:41
 * @LastEditors: NyanCatda
 * @Description: 文件操作
 * @FilePath: \AyaLog\File.go
 */
package AyaLog

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/**
 * @description: 将日志写入日志文件
 * @param {...any} Text 日志内容
 * @return {error} 错误信息
 */
func (Log *Log) writeLogFile(Text ...any) error {
	// 获取一个文件实例
	LogFile, err := Log.openLogFile()
	if err != nil {
		return err
	}
	Write := bufio.NewWriter(LogFile)

	var FileLogText string

	// 遍历消息内容去除颜色
	for _, Value := range Text {
		DelColorText := DelColor(fmt.Sprint(Value))
		FileLogText += DelColorText
		FileLogText += " "
	}

	// 追加打印前缀
	if Log.PrefixWriteFile {
		FileLogText = Log.Prefix + FileLogText
	}

	// 追加打印后缀
	if Log.SuffixWriteFile {
		FileLogText = FileLogText + Log.Suffix
	}

	_, err = Write.WriteString(FileLogText + "\n")
	if err != nil {
		return err
	}
	err = Write.Flush()
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: 打开Log文件，按日志文件分割标识分割日志
 * @param {*}
 * @return {*os.File}
 * @return {error}
 */
func (Log *Log) openLogFile() (*os.File, error) {
	LogFileSegmentation := time.Now().Format(Log.Segmentation)
	// 判断日志分割标识是否变更，变更则关闭文件后重新打开，否则直接返回
	if Log.nowSegmentation != LogFileSegmentation {
		// 关闭先前的文件实例
		if Log.logFile != nil {
			err := Log.logFile.Close()
			if err != nil {
				return nil, err
			}
		}

		// 打开新的文件实例
		Logfile, err := Log.newLogFile(LogFileSegmentation)
		if err != nil {
			return nil, err
		}

		// 将文件实例缓存
		Log.logFile = Logfile
		// 将日志分割标识缓存
		Log.nowSegmentation = LogFileSegmentation

		return Logfile, nil
	} else {
		return Log.logFile, nil
	}
}

/**
 * @description: 打开一个新的日志文件实例
 * @param {string} FileSegmentation 日志文件分割标识
 * @return {*os.File} 日志文件实例
 * @return {error} 错误信息
 */
func (Log *Log) newLogFile(FileSegmentation string) (*os.File, error) {
	// 判断文件夹是否存在，不存在则创建
	mkDir(Log.Path)
	LogFileName := FileSegmentation + ".log"

	LogFilePath := filepath.Clean(Log.Path + "/" + LogFileName)
	Logfile, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return Logfile, nil
}

/**
 * @description: 创建文件夹，如果不存在则创建
 * @param {string} path 文件夹路径
 * @return {bool} 是否创建成功
 * @return {error} 错误信息
 */
func mkDir(Path string) (bool, error) {
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
