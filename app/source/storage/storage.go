package storage

import (
	redisgo "github.com/Conty111/AvitoTech/storage/redisGo"
	"go.uber.org/zap"
)

type Storage interface {
	SetValue(key, value string) error
	GetValue(key string) (string, error)
	Delete(key []string) error
	Connect() error
	Disconnect() error
}

// Возвращает клиент БД, реализующий Storage
func New(addr, pass string, port, DBnum int, logger *zap.Logger) Storage {
	db := redisgo.NewRedis(addr, pass, port, DBnum, logger)
	return db
}
