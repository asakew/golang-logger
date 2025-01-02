package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger.Info("This is an info message.")
}

// {"level":"info","ts":1735780408.776327,"caller":"framework-zap/main.go:8","msg":"This is an info message."}
