<!--
 * @Author: NyanCatda
 * @Date: 2022-05-22 22:28:05
 * @LastEditTime: 2023-01-07 22:49:33
 * @LastEditors: NyanCatda
 * @Description: 自述文件
 * @FilePath: \AyaLog\README.md
-->
# AyaLog
适用于Web后端的Golang Log库

使用简单，适配Gin，Gorm，实现了基础的Log功能，例如日志级别，按时间分割日志与中间件支持，适用于轻量的Log记录需求

# 🎬 如何使用
## 安装
```
go get -u github.com/nyancatda/AyaLog/v2
```

## 基础功能
### 例子: 
``` go
package main

import (
	"errors"

	"github.com/nyancatda/AyaLog/v2"
)

func main() {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()
	// 配置日志实例
	Log.Level = AyaLog.DEBUG // 设置日志等级

	// 打印DeBug日志
	Log.DeBug("System", "This is a debug message") // 2022-05-22 23:57:38 DEBUG [System] This is a debug message
	// 打印Info日志
	Log.Info("System", "This is a info message") // 2022-05-22 23:57:38 INFO [System] This is a info message
	// 打印Warning日志
	Log.Warning("System", "This is a warning message") // 2022-05-22 23:57:38 WARNING [System] This is a warning message
	// 打印Error日志
	Log.Error("System", errors.New("This is a error message")) // 2022-05-22 23:57:38 ERROR [System] This is a error message

	// 为打印的文本设置颜色
	Log.Info("System", "This is "+AyaLog.Green("Green"))
	// 为打印的文本设置背景颜色
	Log.Info("System", "This is "+AyaLog.GreenBackground("GreenBackground"))
}
```

## 为Gin日志启用
### 安装Gin日志模块
```
go get -u github.com/nyancatda/AyaLog/Module/GinLog
```
### 注册模块提供的日志中间件
``` go
r.Use((GinLog.GinLog(*Log)))
```
### 例子
``` go
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog/Module/GinLog"
	"github.com/nyancatda/AyaLog/v2"
)

func main() {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()

	// 关闭Gin默认的日志输出
	gin.DefaultWriter = os.Stdin
	// 初始化GIN
	r := gin.Default()
	// 注册日志中间件
	r.Use((GinLog.GinLog(*Log)))

	// 运行
	if err := r.Run(":8000"); err != nil {
		Log.Error("GIN", err)
	}
}
```

## 为Gorm日志启用
### 安装Gorm日志模块
```
go get -u github.com/nyancatda/AyaLog/Module/GormLog
```
### 将Logger设置为模块提供的接口
``` go
ConnectDB, err := gorm.Open(mysql.Open(ConnectInfo), &gorm.Config{
	Logger: GormLog.GormLog{Log: *Log},
})
```
### 例子
``` go
package main

import (
	"github.com/nyancatda/AyaLog/Module/GormLog"
	"github.com/nyancatda/AyaLog/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()

	// 配置数据库信息
	MySQLUser := ""
	MySQLPassword := ""
	MySQLIp := ""
	MySQLDatabase := ""

	ConnectInfo := MySQLUser + ":" + MySQLPassword + "@tcp(" + MySQLIp + ")/" + MySQLDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	//创建MySQL连接
	ConnectDB, err := gorm.Open(mysql.Open(ConnectInfo), &gorm.Config{
		Logger: GormLog.GormLog{Log: *Log}, // Logger设置为AyaLog的GormLog模块
	})
	if err != nil {
		Log.Error("Gorm", err)
	}

	// 关闭连接
	SQLDB, err := ConnectDB.DB()
	if err != nil {
		Log.Error("Gorm", err)
	}
	defer SQLDB.Close()
}
```

## 启用自动压缩与清理日志文件
### 安装定时任务模块
```
go get -u github.com/nyancatda/AyaLog/Module/TimedTask
```
### 直接启用定时任务
直接启动默认的定时任务，每天压缩日志文件，每天清理7天前的日志文件
``` go
package main

import (
	"github.com/nyancatda/AyaLog/Module/TimedTask"
	"github.com/nyancatda/AyaLog/v2"
)

func main() {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()

	// 启动定时任务
	go TimedTask.Start(*Log)

	Log.Info("System", "定时任务启动")
}
```
### 自定义定时任务
使用模块提供的函数，自定义定时任务，推荐使用`jasonlvhit/gocron`
``` go
package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/nyancatda/AyaLog/Module/TimedTask"
	"github.com/nyancatda/AyaLog/v2"
)

func main() {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()

	// 新建一个线程来执行定时任务
	go func() {
		// 初始化定时任务
		Task := gocron.NewScheduler()

		Task.Every(1).Day().Do(TimedTask.CompressLogs, *Log) // 每天执行一次日志压缩任务
		Task.Every(1).Day().Do(TimedTask.CleanFile, *Log, 7) // 每天执行一次日志清理任务，清理7天前的日志文件

		// 开始执行定时任务
		<-Task.Start()
	}()

	Log.Info("System", "定时任务启动")
}
```

## 使用中间件
AyaLog提供了中间件的支持，可以方便的在日志打印前后进行调整和上报等操作
### 打印前中间件
打印前中间件会在Print函数的最开始执行，即使使用OFF日志等级也会执行，且无视日志等级设定
``` go
Log.UseBefore(func(Level *int, Source *string, Text ...*any) {
	fmt.Println("I will print before")

	// 可以使用指针修改日志内容
	*Source = "Middleware"
})
```
### 打印后中间件
打印后中间件会在Print函数末尾执行，内容会受到打印前中间件的影响，且无法修改日志内容
``` go
Log.UseAfter(func(Level int, Source string, Text ...any) {
	fmt.Println("I will print after")
})
```
### 完整示例
``` go
package main

import (
	"fmt"

	"github.com/nyancatda/AyaLog/v2"
)

func main() {
	// 创建一个默认日志实例
	Log := AyaLog.NewLog()

	// 添加打印前中间件
	Log.UseBefore(func(Level *int, Source *string, Text ...*any) {
		fmt.Println("I will print before")

		*Source = "Middleware"
	})

	// 添加打印后中间件
	Log.UseAfter(func(Level int, Source string, Text ...any) {
		fmt.Println("I will print after")
	})

	// 打印Info日志
	Log.Info("System", "This is a info message")
}
```
此示例会输出
```
I will print before
2022-05-22 23:57:38 INFO [Middleware] This is a info message
I will print after
```

# 📖 许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档