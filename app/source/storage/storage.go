package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Storage interface {
	SetValue(key, value string)
	GetValue(key string) string
	Delete(key string)
	Connect()
	Disconnect()
}

type RedisDB struct {
	Addr     string
	Password string
	DBnum    int
	Client   *redis.Client
}

func NewRedis(addr, pass string, DBnum int) *RedisDB {
	if DBnum < 0 {
		DBnum = 0
	}

	return &RedisDB{
		Addr:     addr,
		Password: pass,
		DBnum:    DBnum,
	}
}

func (db *RedisDB) Connect() {
	db.Client = redis.NewClient(&redis.Options{
		Addr:     db.Addr,
		Password: db.Password,
		DB:       db.DBnum,
		// TLSConfig: &tls.Config{}, // Здесь добавлять сертификаты
	})
}

func (db *RedisDB) Disconnect() {

}

func (db *RedisDB) SetValue(key, value string) {
	err := db.Client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}

func (db *RedisDB) GetValue(key string) string {
	val, err := db.Client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
func (db *RedisDB) Delete(keys []string) {
	// status, err := db.Client.Del(keys...).Result()
}
