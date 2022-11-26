/*
 * @Author: NyanCatda
 * @Date: 2022-06-23 22:18:07
 * @LastEditTime: 2022-11-26 23:43:00
 * @LastEditors: NyanCatda
 * @Description: 定时任务模块
 * @FilePath: \TimedTask\TimedTask.go
 */
package TimedTask

import (
	"github.com/jasonlvhit/gocron"
	"github.com/nyancatda/AyaLog/v2"
)

/**
 * @description: 启动定时任务
 * @param {AyaLog.Log} Log 日志实例
 * @return {*}
 */
func Start(Log AyaLog.Log) {
	// 初始化定时任务
	Task := gocron.NewScheduler()

	Task.Every(1).Day().Do(CompressLogs, Log) // 每天执行一次日志压缩任务
	Task.Every(1).Day().Do(CleanFile, Log, 7) // 每天执行一次日志清理任务，清理7天前的日志文件

	// 开始执行定时任务
	<-Task.Start()
}
