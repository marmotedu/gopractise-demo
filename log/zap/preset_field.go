package main

import "go.uber.org/zap"

func main() {
	logger := zap.NewExample(zap.Fields(
		zap.Int("userID", 10),
		zap.String("requestID", "fbf54504"),
	))

	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
}
