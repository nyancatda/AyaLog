/*
 * @Author: NyanCatda
 * @Date: 2022-06-23 22:24:03
 * @LastEditTime: 2022-06-23 22:53:09
 * @LastEditors: NyanCatda
 * @Description: 定时压缩文件
 * @FilePath: \TimedTask\CompressLogs.go
 */
package TimedTask

import (
	"compress/gzip"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/nyancatda/AyaLog"
)

/**
 * @description: 压缩日志文件，会压缩除当日以外的所有日志文件
 * @return {*}
 */
func CompressLogs() {
	// 获取当前时间，按照日志文件分割标识
	NowTime := time.Now().Format(AyaLog.LogSegmentation)

	// 遍历Log目录下的所有文件
	FileList, err := GetFileList(AyaLog.LogPath)
	if err != nil {
		AyaLog.Error("CompressLogs", err)
		return
	}

	// 压缩日志文件
	for _, FileName := range FileList {
		// 如果是当日的日志文件，则跳过
		if FileName == NowTime+".log" {
			continue
		}

		// 压缩日志文件
		FilePath := filepath.Clean(AyaLog.LogPath + "/" + FileName)
		err := WriterGZip(FilePath)
		if err != nil {
			AyaLog.Error("CompressLogs", err)
			return
		}

		// 删除源日志文件
		err = os.Remove(FilePath)
		if err != nil {
			AyaLog.Error("CompressLogs", err)
			return
		}
	}
}

/**
 * @description: 遍历获取目录下的所有文件
 * @param {string} FilePath 文件路径
 * @return {[]string} 文件名称数组
 * @return {error} 错误
 */
func GetFileList(FilePath string) ([]string, error) {
	var FileList []string
	Files, err := ioutil.ReadDir(FilePath)
	if err != nil {
		return FileList, err
	}
	for _, File := range Files {
		if !File.IsDir() {
			FileList = append(FileList, File.Name())
		}
	}
	return FileList, err
}

/**
 * @description: 写入GZip压缩文件
 * @param {string} FilePath 文件路径
 * @return {error} 错误
 */
func WriterGZip(FilePath string) error {
	// 创建文件
	ZipFile, err := os.Create(FilePath + ".gzip")
	if err != nil {
		return err
	}
	defer ZipFile.Close()

	GZipFileWriter := gzip.NewWriter(ZipFile)
	defer GZipFileWriter.Close()

	// 获取要打包的文件
	File, err := os.Open(FilePath)
	if err != nil {
		return err
	}
	defer File.Close()

	// 获取要打包的文件信息
	FileInfo, err := File.Stat()
	if err != nil {
		return err
	}

	// 读取文件数据
	Buffer := make([]byte, FileInfo.Size())
	_, err = File.Read(Buffer)
	if err != nil {
		return err
	}

	// 写入数据到GZip包
	GZipFileWriter.Header.Name = FileInfo.Name()
	_, err = GZipFileWriter.Write(Buffer)
	if err != nil {
		return err
	}

	return nil
}
