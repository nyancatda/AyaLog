/*
 * @Author: NyanCatda
 * @Date: 2022-05-23 12:51:36
 * @LastEditTime: 2023-01-07 22:22:57
 * @LastEditors: NyanCatda
 * @Description: 输出测试
 * @FilePath: \AyaLog\Test\Print_test.go
 */
package Test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nyancatda/AyaLog/v2"
)

func TestPrint(t *testing.T) {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()
	// 配置日志实例
	Log.Level = AyaLog.DEBUG // 设置日志等级

	// 设置日志打印前中间件
	Log.UseBefore(func(Level *int, Source *string, Text ...*any) {
		// 修改内容
		*Source = "Middleware"
	})
	// 设置日志打印后中间件
	Log.UseAfter(func(Level int, Source string, Text ...any) {
		fmt.Println("After: ", Level, Source, Text)
	})

	// 打印DeBug日志
	Log.DeBug("System", "This is a debug message")
	// 打印Info日志
	Log.Info("System", "This is a info message")
	// 打印Warning日志
	Log.Warning("System", "This is a warning message")
	// 打印Error日志
	Log.Error("System", errors.New("This is a error message"))

	// 为打印的文本设置颜色
	Log.Info("System", "This is "+AyaLog.Green("Green"))
	// 为打印的文本设置背景颜色
	Log.Info("System", "This is "+AyaLog.GreenBackground("GreenBackground"))
}
