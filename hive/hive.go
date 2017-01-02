package main

import (
	"github.com/go2run/hive/performance"
	"github.com/uber-go/zap"
)

func main() {
	initBee()
}

func initBee() {
	logger := zap.New(
		zap.NewJSONEncoder(zap.NoTime()), // drop timestamps in tests
	)

	logger.Info("Begin to Init hive.")
	performance.Init()
	logger.Info("Init hive successfully.")
}
