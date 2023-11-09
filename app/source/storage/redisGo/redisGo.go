package redisgo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

type RedisDB struct {
	Addr     string
	Password string
	DBnum    int
	Client   *redis.Client
	Logger   *zap.Logger
}

func NewRedis(addr, pass string, port, DBnum int, logger *zap.Logger) *RedisDB {
	if DBnum < 0 {
		DBnum = 0
	}
	addr += fmt.Sprintf(":%d", port)
	db := &RedisDB{
		Addr:     addr,
		Password: pass,
		DBnum:    DBnum,
		Logger:   logger,
	}
	err := db.Connect()
	if err != nil {
		db.Logger.Fatal("Cannot connect to the database", zap.Error(err))
	}
	db.Logger.Info("Connected to the database")
	return db
}

func (db *RedisDB) Connect() error {
	db.Client = redis.NewClient(&redis.Options{
		Addr:     db.Addr,
		Password: db.Password,
		DB:       db.DBnum,
		// TLSConfig: &tls.Config{}, // Здесь добавлять сертификаты
	})
	return nil
}

func (db *RedisDB) Disconnect() error {
	db.Logger.Warn("Disconnecting from database")
	err := db.Client.Conn().Close()
	if err == nil {
		return err
	}
	return nil
}

func (db *RedisDB) SetValue(key, value string) error {
	res, err := db.Client.Set(ctx, key, value, 0).Result()
	db.Logger.Info("Setting value from DB", zap.String("status", res))
	if err != nil {
		return err
	}
	return nil
}

func (db *RedisDB) GetValue(key string) (string, error) {
	val, err := db.Client.Get(ctx, key).Result()
	db.Logger.Info("Getting value from DB", zap.String("key", key))
	if errors.Is(err, redis.Nil) {
		return "", sql.ErrNoRows
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (db *RedisDB) Delete(keys []string) error {
	status, err := db.Client.Del(ctx, keys...).Result()
	db.Logger.Info("Deleting key-value from DB", zap.Int("status", int(status)))
	if err != nil {
		return err
	}
	return nil
}
