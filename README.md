<!--
 * @Author: NyanCatda
 * @Date: 2022-05-22 22:28:05
 * @LastEditTime: 2022-06-23 23:35:38
 * @LastEditors: NyanCatda
 * @Description: 自述文件
 * @FilePath: \AyaLog\README.md
-->
# AyaLog
适用于Web后端的Golang Log库

使用简单，适配Gin，Gorm，实现了基础的Log功能，例如日志级别，按时间分割日志，适用于轻量的Log记录需求

# 🎬如何使用
## 安装
```
go get -u github.com/nyancatda/AyaLog
```

## 基础功能
### 例子: 
``` go
package main

import (
	"errors"

	"github.com/nyancatda/AyaLog"
)

func main() {
	// 设置Log参数
	AyaLog.LogLevel = AyaLog.DEBUG        // 设置Log等级
	AyaLog.LogPath = "./logs/"            // 设置Log路径
	AyaLog.LogSegmentation = "2006-01-02" // 设置Log分割标识，使用Go的时间格式化标识
	AyaLog.LogWriteFile = true            // 设置是否写入文件
	AyaLog.ColorPrint = true              // 设置是否打印颜色

	// 打印DeBug日志
	AyaLog.DeBug("System", "This is a debug message") // 2022-05-22 23:57:38 DEBUG [System] This is a debug message
	// 打印Info日志
	AyaLog.Info("System", "This is a info message") // 2022-05-22 23:57:38 INFO [System] This is a info message
	// 打印Warning日志
	AyaLog.Warning("System", "This is a warning message") // 2022-05-22 23:57:38 WARNING [System] This is a warning message
	// 打印Error日志
	AyaLog.Error("System", errors.New("This is a error message")) // 2022-05-22 23:57:38 ERROR [System] This is a error message

	// 为打印的文本设置颜色
	AyaLog.Info("System", "This is "+AyaLog.Green("Green"))
	// 为打印的文本设置背景颜色
	AyaLog.Info("System", "This is "+AyaLog.GreenBackground("GreenBackground"))
}
```

## 为Gin日志启用
### 安装Gin日志模块
```
go get -u github.com/nyancatda/AyaLog/ModLog/GinLog
```
### 注册模块提供的日志中间件
``` go
r.Use((GinLog.GinLog()))
```
### 例子
``` go
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/AyaLog/ModLog/GinLog"
)

func main() {
	// 关闭Gin默认的日志输出
	gin.DefaultWriter = ioutil.Discard
	// 初始化GIN
	r := gin.Default()
	// 注册日志中间件
	r.Use((GinLog.GinLog()))

	// 运行
	if err := r.Run(":8000"); err != nil {
		AyaLog.Error("System", err)
	}
}
```

## 为Gorm日志启用
### 安装Gorm日志模块
```
go get -u github.com/nyancatda/AyaLog/ModLog/GormLog
```
### 将Logger设置为模块提供的接口
``` go
ConnectDB, err := gorm.Open(mysql.Open(ConnectInfo), &gorm.Config{
    Logger: GormLog.GormLog{},
})
```
### 例子
``` go
package main

import (
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/AyaLog/ModLog/GormLog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 配置数据库信息
	MySQLUser := ""
	MySQLPassword := ""
	MySQLIp := ""
	MySQLDatabase := ""

	ConnectInfo := MySQLUser + ":" + MySQLPassword + "@tcp(" + MySQLIp + ")/" + MySQLDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	//创建MySQL连接
	ConnectDB, err := gorm.Open(mysql.Open(ConnectInfo), &gorm.Config{
		Logger: GormLog.GormLog{}, // Logger设置为AyaLog的GormLog模块
	})
	if err != nil {
		AyaLog.Error("System", err)
	}

	// 关闭连接
	SQLDB, err := ConnectDB.DB()
	if err != nil {
		AyaLog.Error("System", err)
	}
	defer SQLDB.Close()
}
```

## 启用自动压缩与清理日志文件
### 安装定时任务模块
```
go get -u github.com/nyancatda/AyaLog/TimedTask
```
### 直接启用定时任务
直接启动默认的定时任务，每天压缩日志文件，每天清理7天前的日志文件
``` go
package main

import (
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/AyaLog/TimedTask"
)

func main() {
	// 设置Log参数
	AyaLog.LogLevel = AyaLog.DEBUG        // 设置Log等级
	AyaLog.LogPath = "./logs/"            // 设置Log路径
	AyaLog.LogSegmentation = "2006-01-02" // 设置Log分割标识

	// 启动定时任务
	go TimedTask.Start()

	AyaLog.Info("System", "定时任务启动")
}

```
### 自定义定时任务
使用模块提供的函数，自定义定时任务，推荐使用`jasonlvhit/gocron`
``` go
package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/AyaLog/TimedTask"
)

func main() {
	// 设置Log参数
	AyaLog.LogLevel = AyaLog.DEBUG        // 设置Log等级
	AyaLog.LogPath = "./logs/"            // 设置Log路径
	AyaLog.LogSegmentation = "2006-01-02" // 设置Log分割标识

	// 新建一个线程来执行定时任务
	go func() {
		// 初始化定时任务
		Task := gocron.NewScheduler()

		Task.Every(1).Day().Do(TimedTask.CompressLogs) // 每天执行一次日志压缩任务
		Task.Every(1).Day().Do(TimedTask.CleanFile, 7) // 每天执行一次日志清理任务，清理7天前的日志文件

		// 开始执行定时任务
		<-Task.Start()
	}()

	AyaLog.Info("System", "定时任务启动")
}
```

# 📖许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档