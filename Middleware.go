/*
 * @Author: NyanCatda
 * @Date: 2023-01-07 21:56:56
 * @LastEditTime: 2023-01-07 22:03:16
 * @LastEditors: NyanCatda
 * @Description: 中间件实现
 * @FilePath: \AyaLog\Middleware.go
 */
package AyaLog

type middlewareBefore func(Level *int, Source *string, Text ...*any)
type middlewareAfter func(Level int, Source string, Text ...any)

/**
 * @description: 添加日志打印前中间件
 * @param {*int} Func Level 日志等级 DEBUG(0)/INFO(1)/WARNING(2)/ERROR(3)/OFF(4)
 * @param {*string} Func Source 日志来源
 * @param {...*any} Func Text 日志内容
 * @return {*}
 */
func (Log *Log) UseBefore(Func func(Level *int, Source *string, Text ...*any)) {
	Log.middlewareBefore = append(Log.middlewareBefore, Func)
}

/**
 * @description: 添加日志打印后中间件
 * @param {int} Func Level 日志等级 DEBUG(0)/INFO(1)/WARNING(2)/ERROR(3)/OFF(4)
 * @param {string} Func Source 日志来源
 * @param {...any} Func Text 日志内容
 * @return {*}
 */
func (Log *Log) UseAfter(Func func(Level int, Source string, Text ...any)) {
	Log.middlewareAfter = append(Log.middlewareAfter, Func)
}
