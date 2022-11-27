/*
 * @Author: NyanCatda
 * @Date: 2022-05-22 00:11:00
 * @LastEditTime: 2022-11-26 23:28:49
 * @LastEditors: NyanCatda
 * @Description: Gin日志模块
 * @FilePath: \GinLog\GinLog.go
 */
package GinLog

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog/v2"
)

/**
 * @description: Gin日志中间件
 * @param {AyaLog.Log} Log 日志实例
 * @return {*}
 */
func GinLog(Log AyaLog.LogPrint) gin.HandlerFunc {
	return func(c *gin.Context) {
		StartTime := time.Now()               // 记录开始时间
		c.Next()                              // 继续处理请求
		EndTime := time.Now()                 // 记录结束时间
		LatencyTime := EndTime.Sub(StartTime) // 执行时间
		ReqMethod := c.Request.Method         // 请求方式
		ReqUri := c.Request.RequestURI        // 请求路由
		StatusCode := c.Writer.Status()       // 状态码
		ClientIP := c.ClientIP()              // 请求IP

		// 状态码添加颜色
		var StatusCodeText string
		switch StatusCode / 100 % 10 {
		case 2:
			StatusCodeText = AyaLog.GreenBackground(fmt.Sprintf(" %3d ", StatusCode))
		case 3:
			StatusCodeText = AyaLog.WhiteBackground(fmt.Sprintf(" %3d ", StatusCode))
		case 4:
			StatusCodeText = AyaLog.YellowBackground(fmt.Sprintf(" %3d ", StatusCode))
		default:
			StatusCodeText = AyaLog.RedBackground(fmt.Sprintf(" %3d ", StatusCode))
		}

		// 请求方式添加颜色
		ReqMethodText := " " + ReqMethod + "      "
		switch ReqMethod {
		case "GET":
			ReqMethodText = AyaLog.BlueBackground(ReqMethodText)
		case "POST":
			ReqMethodText = AyaLog.CyanBackground(ReqMethodText)
		case "PUT":
			ReqMethodText = AyaLog.YellowBackground(ReqMethodText)
		case "PATCH":
			ReqMethodText = AyaLog.GreenBackground(ReqMethodText)
		case "DELETE":
			ReqMethodText = AyaLog.RedBackground(ReqMethodText)
		case "HEAD":
			ReqMethodText = AyaLog.MagentaBackground(ReqMethodText)
		case "OPTIONS":
			ReqMethodText = AyaLog.WhiteBackground(ReqMethodText)
		}

		// 日志格式
		LogInfo := fmt.Sprintf("%s| %13v | %15s |%s\"%s\"",
			StatusCodeText,
			LatencyTime,
			ClientIP,
			ReqMethodText,
			ReqUri,
		)

		// 打印日志
		Log.Info("GIN", LogInfo)
	}
}
