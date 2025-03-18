# utils
Standard library extensions

各种工具库

## scheduler
### engine
一个任务调度框架，可以控制goroutine数量,任务失败重试，任务衍生子任务执行，任务检测，任务统计
### crawler
爬虫框架，基于scheduler/engine

## dao
各种常用dao的封装
### database
数据库操作封装,主要针对gorm扩展开发
## eflag
通过环境变量 flag 注入结构体
## log
zap的二次封装,开箱即用
## iter
标准库iter的扩展,stream操作实现
## net
