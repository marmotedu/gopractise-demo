package main

import "go.uber.org/zap"

func main() {
	zap.L().Info("default global Logger")
	zap.S().Info("default global SugaredLogger")

	logger := zap.NewExample()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	zap.L().Info("replaced global Logger")
	zap.S().Info("replaced global SugaredLogger")
}
