/*
 * @Author: NyanCatda
 * @Date: 2022-05-22 00:02:30
 * @LastEditTime: 2022-05-22 22:41:22
 * @LastEditors: NyanCatda
 * @Description: 终端输出增加颜色
 * @FilePath: \AyaLog\Color.go
 */
package AyaLog

import (
	"fmt"
	"regexp"
)

// 定义文字颜色
const (
	TextWhite = 1
)
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextGrey
)

/**
 * @description: 设置文字颜色为黑色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Black(msg string) string {
	return SetColor(msg, 0, 0, TextBlack)
}

/**
 * @description: 设置文字颜色为红色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Red(msg string) string {
	return SetColor(msg, 0, 0, TextRed)
}

/**
 * @description: 设置文字颜色为绿色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Green(msg string) string {
	return SetColor(msg, 0, 0, TextGreen)
}

/**
 * @description: 设置文字颜色为黄色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextYellow)
}

/**
 * @description: 设置文字颜色为蓝色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextBlue)
}

/**
 * @description: 设置文字颜色为紫红色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextMagenta)
}

/**
 * @description: 设置文字颜色为青蓝色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextCyan)
}

/**
 * @description: 设置文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func White(msg string) string {
	return SetColor(msg, 0, 0, TextWhite)
}

// 定义背景颜色
const (
	BackgroundBlack = iota + 40
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

/**
 * @description: 设置背景颜色为黑色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func BlackBackground(msg string) string {
	return SetColor(msg, 0, BackgroundBlack, TextWhite)
}

/**
 * @description: 设置背景颜色为红色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func RedBackground(msg string) string {
	return SetColor(msg, 0, BackgroundRed, TextWhite)
}

/**
 * @description: 设置背景颜色为绿色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func GreenBackground(msg string) string {
	return SetColor(msg, 0, BackgroundGreen, TextWhite)
}

/**
 * @description: 设置背景颜色为黄色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func YellowBackground(msg string) string {
	return SetColor(msg, 0, BackgroundYellow, TextWhite)
}

/**
 * @description: 设置背景颜色为蓝色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func BlueBackground(msg string) string {
	return SetColor(msg, 0, BackgroundBlue, TextWhite)
}

/**
 * @description: 设置背景颜色为紫红色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func MagentaBackground(msg string) string {
	return SetColor(msg, 0, BackgroundMagenta, TextWhite)
}

/**
 * @description: 设置背景颜色为青蓝色，文字颜色为白色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func CyanBackground(msg string) string {
	return SetColor(msg, 0, BackgroundCyan, TextWhite)
}

/**
 * @description: 设置背景颜色为白色，文字颜色为黑色
 * @param {string} msg 需要设置颜色的文字
 * @return {string} 设置颜色后的文字
 */
func WhiteBackground(msg string) string {
	return SetColor(msg, 0, BackgroundWhite, TextBlack)
}

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  灰色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

/**
 * @description: 设置文字颜色与背景颜色
 * @param {string} msg 需要设置颜色的文字
 * @param {int} conf 颜色配置
 * @param {int} bg 设置背景颜色
 * @param {int} text 设置文字颜色
 * @return {string} 设置颜色后的文字
 */
func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}

/**
 * @description: 去除文字颜色
 * @param {string} msg 需要去除颜色的文字
 * @return {string} 去除颜色后的文字
 */
func DelColor(msg string) string {
	reg := regexp.MustCompile(`\x1b(\[.*?[@-~]|\].*?(\x07|\x1b\\))`)
	return reg.ReplaceAllString(msg, "")
}
