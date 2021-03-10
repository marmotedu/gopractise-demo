package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	url := "http://marmotedu.com"
	// 结构化日志
	logger.Sugar().Infow("failed to fetch URL", "url", url, "attempt", 3, "backoff", time.Second)

	// 非结构化日志
	logger.Sugar().Infof("failed to fetch URL: %s", url)
}
