package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Настраиваем конфигурацию логгера
func NewLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig() // Уровень логирования
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	logger, err := config.Build() // Настраиваем логгер с конфигом
	if err != nil {
		fmt.Printf("Ошибка настройки логгера: %v\n", err)
	}
	return logger
}
