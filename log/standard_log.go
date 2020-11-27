package main

import (
	"log"
	"os"
)

func main() {
	// 输出到文件
	logFile, err := os.Create("./log.log")
	defer logFile.Close()
	if err != nil {
		log.Fatalln("create file log.log failed")
	}
	logger := log.New(logFile, "[Debug] ", log.Lshortfile)
	logger.SetOutput(os.Stdout)
	logger.Print("call Print: line1")
	logger.Println("call Println: line2")

	// 修改日志配置
	logger.SetPrefix("[Info] ")
	logger.SetFlags(log.Ldate)
	logger.SetOutput(os.Stdout)
	logger.Print("Info check stdout")
}
