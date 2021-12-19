package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func helloCron() {
	fmt.Println("hello cron")
}

func main() {
	fmt.Println("starting go cron...")

	// 创建一个cron实例
	cron := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(nil), cron.Recover(nil)))

	// 添加一个定时任务
	cron.AddFunc("*  *  *  *  *  *", helloCron)

	// 启动计划任务
	cron.Start()

	// 关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer cron.Stop()

	select {} // 查询语句，保持程序运行，在这里等同于for{}
}
