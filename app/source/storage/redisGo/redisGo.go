package redisgo

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisDB struct {
	Addr     string
	Password string
	DBnum    int
	Client   *redis.Client
}

func NewRedis(addr, pass string, port, DBnum int) *RedisDB {
	if DBnum < 0 {
		DBnum = 0
	}
	addr += fmt.Sprintf(":%d", port)
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
	log.Print("Connecting to the database")
}

func (db *RedisDB) Disconnect() {
	log.Print("Disconnecting")
	db.Client.Conn().Close()
}

func (db *RedisDB) SetValue(key, value string) {
	res, err := db.Client.Set(ctx, key, value, 0).Result()
	log.Print("Setting value status:", res)
	if err != nil {
		log.Panic(err)
	}
}

func (db *RedisDB) GetValue(key string) string {
	val, err := db.Client.Get(ctx, key).Result()
	if err != nil {
		log.Panic(err)
	}
	return val
}

func (db *RedisDB) Delete(keys []string) {
	status, err := db.Client.Del(ctx, keys...).Result()
	log.Print("Deleting status:", status)
	if err != nil {
		log.Panic(err)
	}
}
