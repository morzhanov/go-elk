package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	baseLogsPath := "./log/go-elk-example/"
	if err := os.MkdirAll(baseLogsPath, 0777); err != nil {
		return nil, err
	}
	filePath := fmt.Sprintf("%sapp.log", baseLogsPath)
	_, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	cfg.OutputPaths = []string{filePath, "stdout"}
	return cfg.Build()
}
