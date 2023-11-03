package storage

import redisgo "github.com/Conty111/AvitoTech/storage/redisGo"

type Storage interface {
	SetValue(key, value string)
	GetValue(key string) string
	Delete(key []string)
	Connect()
	Disconnect()
}

func New(addr, pass string, port, DBnum int) Storage {
	// Возвращает клиент БД, реализующий Storage
	return redisgo.NewRedis(addr, pass, port, DBnum)
}
