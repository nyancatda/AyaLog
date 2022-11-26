/*
 * @Author: NyanCatda
 * @Date: 2022-11-26 16:45:50
 * @LastEditTime: 2022-11-26 18:05:11
 * @LastEditors: NyanCatda
 * @Description: 文件操作
 * @FilePath: \AyaLog\File.go
 */
package AyaLog

import (
	"os"
	"path/filepath"
	"time"
)

/**
 * @description: 打开Log文件，按日志文件分割标识分割日志
 * @param {*}
 * @return {*os.File}
 * @return {error}
 */
func (Log Log) openLogFile() (*os.File, error) {
	// 判断文件夹是否存在，不存在则创建
	mkDir(Log.Path)

	LogFileName := time.Now().Format(Log.Segmentation) + ".log"

	Logfile, err := os.OpenFile(Log.Path+LogFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
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
