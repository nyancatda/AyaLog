/*
 * @Author: NyanCatda
 * @Date: 2022-05-23 12:51:36
 * @LastEditTime: 2022-05-23 13:01:06
 * @LastEditors: NyanCatda
 * @Description: 输出测试
 * @FilePath: \AyaLog\Test\Print_test.go
 */
package Test

import (
	"errors"
	"testing"

	"github.com/nyancatda/AyaLog"
)

func TestPrint(t *testing.T) {
	// 设置Log参数
	AyaLog.LogLevel = AyaLog.OFF // 设置Log等级
	AyaLog.LogPath = "../logs/"  // 设置Log路径
	AyaLog.ColorPrint = true     // 设置是否打印颜色

	// 打印DeBug日志
	AyaLog.DeBug("System", "This is a debug message")
	// 打印Info日志
	AyaLog.Info("System", "This is a info message")
	// 打印Warning日志
	AyaLog.Warning("System", "This is a warning message")
	// 打印Error日志
	AyaLog.Error("System", errors.New("This is a error message"))

	// 为打印的文本设置颜色
	AyaLog.Info("System", "This is "+AyaLog.Green("Green"))
	// 为打印的文本设置背景颜色
	AyaLog.Info("System", "This is "+AyaLog.GreenBackground("GreenBackground"))
}
