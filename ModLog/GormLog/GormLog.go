/*
 * @Author: NyanCatda
 * @Date: 2022-05-22 02:31:00
 * @LastEditTime: 2022-05-22 22:43:49
 * @LastEditors: NyanCatda
 * @Description: Gorm日志模块
 * @FilePath: \AyaLog\ModLog\GormLog\GormLog.go
 */
package GormLog

import (
	"context"
	"fmt"
	"time"

	"github.com/nyancatda/AyaLog"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLog struct {
}

/**
 * @description: 日志模型(废弃函数)
 * @param {logger.LogLevel} LogLevel 日志等级
 * @return {*}
 */
func (g GormLog) LogMode(logger.LogLevel) logger.Interface {
	return GormLog{}
}

/**
 * @description: 打印Info日志
 * @param {context.Context} ctx 上下文
 * @param {string} msg 日志内容
 * @param {...interface{}} data 日志内容
 * @return {*}
 */
func (g GormLog) Info(ctx context.Context, msg string, data ...interface{}) {
	LogText := fmt.Sprintf("%s "+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	AyaLog.Info("Gorm", LogText)
}

/**
 * @description: 打印Warning日志
 * @param {context.Context} ctx 上下文
 * @param {string} msg 日志内容
 * @param {...interface{}} data 日志内容
 * @return {*}
 */
func (g GormLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	LogText := fmt.Sprintf("%s "+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	AyaLog.Warning("Gorm", LogText)
}

/**
 * @description: 打印Error日志
 * @param {context.Context} ctx 上下文
 * @param {string} msg 日志内容
 * @param {...interface{}} data 日志内容
 * @return {*}
 */
func (g GormLog) Error(ctx context.Context, msg string, data ...interface{}) {
	LogText := fmt.Sprintf("%s "+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	AyaLog.Print("Gorm", AyaLog.ERROR, LogText)
}

/**
 * @description: 打印Trace日志
 * @param {context.Context} ctx 上下文
 * @param {time.Time} begin 开始时间
 * @param {func() (string, int64)} fc 函数
 * @param {error} err 错误
 * @return {*}
 */
func (g GormLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		AyaLog.Error("Gorm", err)
	}

	var LogText string
	if rows == -1 {
		LogText = fmt.Sprintf("%.3fms | rows:%v | %s", float64(elapsed.Nanoseconds())/1e6, "-", sql)
	} else {
		LogText = fmt.Sprintf("%.3fms | rows:%v | %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
	}
	AyaLog.Info("Gorm", LogText)
}
