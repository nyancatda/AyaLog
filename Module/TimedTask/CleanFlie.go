/*
 * @Author: NyanCatda
 * @Date: 2022-06-23 22:55:55
 * @LastEditTime: 2022-11-26 23:39:52
 * @LastEditors: NyanCatda
 * @Description: 清理过时的压缩文件
 * @FilePath: \TimedTask\CleanFlie.go
 */
package TimedTask

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nyancatda/AyaLog/v2"
)

/**
 * @description: 清理过时的压缩文件
 * @param {AyaLog.Log} Log 日志实例
 * @param {int} Day 过期天数
 * @return {*}
 */
func CleanFile(Log AyaLog.Log, Day int) {
	// 获取当前时间戳
	NowTime := time.Now().Unix()

	// 遍历Log目录下的所有文件
	FileList, err := GetFileList(Log.Path)
	if err != nil {
		Log.Error("CleanFile", err)
		return
	}

	// 删除过期的压缩文件
	for _, FileName := range FileList {
		// 判断是否是.gz文件
		if filepath.Ext(FileName) != ".gz" {
			continue
		}

		// 从文件名字获取保存时间
		FileSave := strings.Replace(FileName, ".log.gz", "", -1)
		// 按照日志文件分割标识转换为时间戳
		FileSaveTime, err := time.ParseInLocation(Log.Segmentation, FileSave, time.Local)
		if err != nil {
			Log.Error("CleanFile", err)
			return
		}
		FileSaveTimeUnix := FileSaveTime.Unix()

		// 获取过期时间戳
		ExpireTime := NowTime - int64(Day*24*60*60)

		// 如果保存时间戳小于过期时间戳，则删除该文件
		if FileSaveTimeUnix < ExpireTime {
			FilePath := filepath.Clean(Log.Path + "/" + FileName)
			err := os.Remove(FilePath)
			if err != nil {
				Log.Error("CleanFile", err)
				return
			}
		}
	}
}
